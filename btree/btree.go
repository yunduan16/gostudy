package btree

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
)

type Item interface {
	Less(than Item) bool
}

const (
	DefaultFreeListSize = 32
)

var (
	nilItems    = make(items, 16)
	nilChildren = make(children, 16)
)

type FreeList struct {
	mu       sync.Mutex
	freeList []*node
}

type children []*node
type freeType int

const (
	ftFreelistFull freeType = iota //node was freed(avaliable for GC, not stored in freelist)
	ftStored                       //node was stored in the freelist for later use
	ftNotOwned                     //node was ignored by COW, since it's owned by another one
)

type items []Item
type Int int

type copyOnWriteContext struct {
	freelist *FreeList
}

//以下是 copyOnWriteContext一些方法
func (c *copyOnWriteContext) newNode() (n *node) {
	n = c.freelist.newNode()
	n.cow = c
	return
}

//
func (c *copyOnWriteContext) freeNode(n *node) freeType {
	if n.cow == c {
		//clear to allow GC
		n.items.truncate(0)
		n.children.truncate(0)
		n.cow = nil
		//如果空闲队列没满，可以暂存
		if c.freelist.freeNode(n) {
			return ftStored
		} else {
			return ftFreelistFull
		}
	} else {
		return ftNotOwned
	}
}

// It must at all times maintain the invariant that either
//   * len(children) == 0, len(items) unconstrained
//   * len(children) == len(items) + 1
//如果是叶子结点 len(children) = 0 ,包含的关键词len(items)(存的是该关键词下所有对应的数据)不限制
//如果是非叶子结点, len(children) == len(items)+1 ,即子树个数=关键字(存索引地方，一般是items中第一个元素)+1

type node struct {
	items    items
	children children
	cow      *copyOnWriteContext
}

type BTree struct {
	degree int //度，拥有子结点的个数，一般只最小度数
	length int //子结点个数
	root   *node
	cow    *copyOnWriteContext
}

type toRemove int

const (
	removeItem toRemove = iota
	removeMin
	removeMax
)

//定义方向
type direction int

const (
	descend = direction(-1) //下降
	ascend  = direction(+1) //上升
)

//实现Item接口 Less
func (a Int) Less(b Item) bool {
	return a < b.(Int)
}

//FreeList构造方法
func NewFreeList(size int) *FreeList {
	return &FreeList{
		freeList: make([]*node, 0, size),
	}
}

//FreeList的新增Node，如果有空闲的list，则从队列末尾删除一个返回
func (f *FreeList) newNode() (n *node) {
	f.mu.Lock()
	index := len(f.freeList) - 1
	if index < 0 {
		f.mu.Unlock()
		return new(node)
	}
	n = f.freeList[index]
	f.freeList[index] = nil
	f.freeList = f.freeList[:index]
	f.mu.Unlock()
	return
}

//FreeList 添加freelist队列的方法
func (f *FreeList) freeNode(n *node) (out bool) {
	f.mu.Lock()
	if len(f.freeList) < cap(f.freeList) {
		f.freeList = append(f.freeList, n)
		//直接retrun out 会导致删除中断
		out = true
	}
	f.mu.Unlock()
	return
}

type ItemIterator func(i Item) bool

//以下是BTree方法
func New(degree int) *BTree {
	return NewWithFreeList(degree, NewFreeList(DefaultFreeListSize))
}

func NewWithFreeList(degree int, f *FreeList) *BTree {
	//degree > 1
	if degree <= 1 {
		panic("bad degree")
	}
	return &BTree{
		degree: degree,
		cow:    &copyOnWriteContext{freelist: f},
	}
}

//以下是node.items一些方法
//在指定位置添加Item
func (s *items) insertAt(index int, item Item) {
	*s = append(*s, nil)
	if index < len(*s) {
		//降index 后面的值后移一位
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = item
}

//在指定位置移出Item
func (s *items) removeAt(index int) (out Item) {
	out = (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return
}

//队列移出最上面一个
func (s *items) pop() (out Item) {
	index := len(*s) - 1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

//truncate 重置index后的数据(包括index)
func (s *items) truncate(index int) {
	var toClear items
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilItems):]
	}
}

func (s items) find(item Item) (index int, found bool) {
	i := sort.Search(len(s), func(i int) bool {
		return item.Less(s[i])
	})
	if i > 0 && !s[i-1].Less(item) {
		return i - 1, true
	}
	return i, false
}

//以下是children的一些方法,与items类似
func (s *children) insertAt(index int, n *node) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = n
}

func (s *children) removeAt(index int) *node {
	n := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return n
}

func (s *children) pop() (out *node) {
	index := len(*s) - 1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

func (s *children) truncate(index int) {
	var toClear children
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilChildren):]
	}
}

//以下是node的方法
//node.原子性操作
func (n *node) mutableFor(cow *copyOnWriteContext) *node {
	//如果给定cow=结点cow。直接返回该结点
	if n.cow == cow {
		return n
	}
	//从cow拿出一个新结点
	out := cow.newNode()
	if cap(out.items) >= len(n.items) {
		out.items = out.items[:len(n.items)]
	} else {
		out.items = make(items, len(n.items), cap(n.items))
	}
	copy(out.items, n.items)
	//Copy chilidre
	if cap(out.children) >= len(n.children) {
		out.children = out.children[:len(n.children)]
	} else {
		out.children = make(children, len(n.children), cap(n.children))
	}
	copy(out.children, n.children)
	return out
}

//扩展node.chilidren
func (n *node) mutableChild(index int) *node {
	c := n.children[index].mutableFor(n.cow)
	n.children[index] = c
	return c
}

//按照index 分割node
func (n *node) Split(index int) (Item, *node) {
	item := n.items[index]
	next := n.cow.freelist.newNode()
	next.items = append(next.items, n.items[index+1:]...)
	n.items.truncate(index)
	if len(n.children) > 0 {
		next.children = append(next.children, n.children[index+1:]...)
		n.children.truncate(index + 1)
	}
	return item, next
}

//检查child 是否可以分割，如果可以就二分为一，返回true
func (n *node) maybeSplitChild(index, maxItems int) bool {
	if len(n.children[index].items) < maxItems {
		return false
	}
	first := n.mutableChild(index)
	item, second := first.Split(maxItems / 2)
	n.items.insertAt(index, item)
	n.children.insertAt(index+1, second)
	return true
}

//插入item，如果有相同的item就替换，如果超过最大maxItems就分割
func (n *node) insert(item Item, maxItems int) Item {
	i, found := n.items.find(item)
	if found {
		out := n.items[i]
		n.items[i] = item
		return out
	}
	if len(n.children) == 0 {
		n.items.insertAt(i, item)
		return nil
	}
	if n.maybeSplitChild(i, maxItems) {
		inTree := n.items[i]
		switch {
		case item.Less(inTree):
		// no change, we want first split node
		case inTree.Less(item):
			i++
		default:
			out := n.items[i]
			n.items[i] = item
			return out
		}
	}
	return n.mutableChild(i).insert(item, maxItems)
}

//获取给定的关键字，递归查找，如果当前没有则寻找子节点
func (n *node) get(key Item) Item {
	i, found := n.items.find(key)
	if found {
		return n.items[i]
	} else if len(n.children) > 0 {
		return n.children[i].get(key)
	}
	return nil
}

//返回第一个子树种最小的items最小的一个
func min(n *node) Item {
	if n == nil {
		return nil
	}
	for len(n.children) > 0 {
		n = n.children[0]
	}
	if len(n.items) == 0 {
		return nil
	}
	return n.items[0]
}

//返回第一个子树种最小的items最小的一个
func max(n *node) Item {
	if n == nil {
		return nil
	}
	for len(n.children) > 0 {
		n = n.children[len(n.children)-1]
	}
	if len(n.items) == 0 {
		return nil
	}
	return n.items[len(n.items)-1]
}

func (n *node) remove(item Item, minItems int, typ toRemove) Item {
	var i int
	var found bool
	switch typ {
	case removeMax:
		if len(n.children) == 0 {
			return n.items.pop()
		}
		i = len(n.items)
	case removeMin:
		if len(n.children) == 0 {
			return n.items.removeAt(0)
		}
		i = 0
	case removeItem:
		i, found = n.items.find(item)
		if len(n.children) == 0 {
			if found {
				return n.items.removeAt(i)
			}
			return nil
		}
	default:
		panic("invalid type")
	}

	//如果有childre
	if len(n.children[i].items) <= minItems {
		return n.growChildAndRemove(i, item, minItems, typ)
	}
	child := n.mutableChild(i)

	if found {
		out := n.items[i]
		n.items[i] = child.remove(nil, minItems, removeMax)
		return out
	}
	return child.remove(item, minItems, typ)
}

func (n *node) growChildAndRemove(i int, item Item, minItems int, typ toRemove) Item {
	if i > 0 && len(n.children[i-1].items) > minItems {
		//steal from left child
		child := n.mutableChild(i)
		stealFrom := n.mutableChild(i - 1)
		stolenItem := stealFrom.items.pop()

		child.items.insertAt(0, n.items[i-1])
		n.items[i-1] = stolenItem
		if len(stealFrom.children) > 0 {
			child.children.insertAt(0, stealFrom.children.pop())
		}
	} else if i < len(n.items) && len(n.children[i+1].items) > minItems {
		//steal from right child
		child := n.mutableChild(i)
		stealFrom := n.mutableChild(i + 1)
		//将右子树第一个元素移到当前结点
		stolenItem := stealFrom.items.removeAt(0)
		child.items = append(child.items, n.items[i])
		n.items[i] = stolenItem
		if len(stealFrom.children) > 0 {
			child.children = append(child.children, stealFrom.children.removeAt(0))
		}
	} else {
		if i >= len(n.items) {
			i--
		}
		child := n.mutableChild(i)
		//merge with right child
		mergeItem := n.items.removeAt(i)
		mergeChild := n.children.removeAt(i + 1)
		child.items = append(child.items, mergeItem)
		child.items = append(child.items, mergeChild.items...)
		child.children = append(child.children, mergeChild.children...)
		n.cow.freeNode(mergeChild)
	}
	return n.remove(item, minItems, typ)
}

//迭代遍历tree种的元素
//指定正序还是倒序
func (n *node) iterate(dir direction, start, stop Item, includeStart bool, hit bool, iter ItemIterator) (bool, bool) {
	var ok, found bool
	var index int
	switch dir {
	//升序
	case ascend:
		//如果指定开始
		if start != nil {
			index, _ = n.items.find(start)
		}
		for i := index; i < len(n.items); i++ {
			if len(n.children) > 0 {
				if hit, ok = n.children[i].iterate(dir, start, stop, includeStart, hit, iter); !ok {
					return hit, false
				}
			}
			//如果start >= items[i] 表示还需要继续查找
			if !includeStart && !hit && start != nil && !start.Less(n.items[i]) {
				hit = true
				continue
			}
			hit = true
			//如果当前值已经>=指定的stop,返回false
			if stop != nil && !n.items[i].Less(stop) {
				return hit, false
			}
			//看是否满足迭代条件函数
			if !iter(n.items[i]) {
				return hit, false
			}
		}

		if len(n.children) > 0 {
			if hit, ok = n.children[len(n.children)-1].iterate(dir, start, stop, includeStart, hit, iter); !ok {
				return hit, false
			}
		}
		//倒序查找
	case descend:
		if start != nil {
			index, found = n.items.find(start)
			if !found {
				index = index - 1
			}
		} else {
			index = len(n.items) - 1
		}
		for i := index; i >= 0; i-- {
			if start != nil && !n.items[i].Less(start) {
				if !includeStart || hit || start.Less(n.items[i]) {
					continue
				}
			}
			if len(n.children) > 0 {
				if hit, ok = n.children[i+1].iterate(dir, start, stop, includeStart, hit, iter); !ok {
					return hit, false
				}
			}
			if stop != nil && !stop.Less(n.items[i]) {
				return hit, false
			}
			hit = true
			if !iter(n.items[i]) {
				return hit, false
			}
		}

		if len(n.children) > 0 {
			if hit, ok = n.children[0].iterate(dir, start, stop, includeStart, hit, iter); !ok {
				return hit, false
			}
		}
	}
	return hit, true
}

//递归打印node，分层的展示
func (n *node) print(w io.Writer, level int) {
	fmt.Fprintf(w, "%sNODE:%v\n", strings.Repeat("  ", level), n.items)
	for _, c := range n.children {
		c.print(w, level+1)
	}
}

//重置copyOnWriteContext，将item都放到对应node的copyOnWriteContext
func (n *node) reset(c *copyOnWriteContext) bool {
	for _, child := range n.children {
		if !child.reset(c) {
			return false
		}
	}
	return c.freeNode(n) != ftFreelistFull
}

//以下是BTree的一些方法，值拷贝

func (t *BTree) s() {
	n := t.root
	var w io.Writer
	level := 0
	for {
		if len(n.children) == 0 && len(n.items) == 0 {
			break
		}
		n.print(w, level+1)
	}
}

func (t *BTree) Clone() (t2 *BTree) {
	cow1, cow2 := *t.cow, *t.cow
	out := *t
	t.cow = &cow1
	out.cow = &cow2
	return &out
}

//获取最大的items
func (t *BTree) maxItems() int {
	return t.degree*2 - 1
}

//获取最小的items
func (t *BTree) minItems() int {
	return t.degree - 1
}

//替换或插入
func (t *BTree) ReplaceOrInsert(item Item) Item {
	if item == nil {
		panic("nil item being added to BTree")
	}
	if t.root == nil {
		t.root = t.cow.newNode()
		t.root.items = append(t.root.items, item)
		t.length++
		return nil
	} else {
		t.root = t.root.mutableFor(t.cow)
		//如果root 大于最大items，需要分裂
		if len(t.root.items) >= t.maxItems() {
			item2, second := t.root.Split(t.maxItems() / 2)
			oldroot := t.root
			t.root = t.cow.newNode()
			t.root.items = append(t.root.items, item2)
			t.root.children = append(t.root.children, oldroot, second)
		}
	}
	out := t.root.insert(item, t.maxItems())
	if out == nil {
		t.length++
	}
	return out
}

func (t *BTree) deleteItem(item Item, typ toRemove) Item {
	if t.root == nil || len(t.root.items) == 0 {
		return nil
	}
	t.root = t.root.mutableFor(t.cow)
	out := t.root.remove(item, t.minItems(), typ)
	if len(t.root.items) == 0 && len(t.root.children) > 0 {
		oldroot := t.root
		//左子树
		t.root = t.root.children[0]
		t.cow.freeNode(oldroot)
	}
	if out != nil {
		t.length--
	}
	return out
}

//删除一个给定的item
func (t *BTree) Delete(item Item) Item {
	return t.deleteItem(item, removeItem)
}

//删除最小
func (t *BTree) DeleteMin() Item {
	return t.deleteItem(nil, removeMin)
}

//删除最大
func (t *BTree) DeleteMax() Item {
	return t.deleteItem(nil, removeMax)
}

//升序方法
//大于等于，小于范围内集合
func (t *BTree) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, greaterOrEqual, lessThan, true, false, iterator)
}

//小于某值的集合
func (t *BTree) AscendLessThan(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, nil, pivot, false, false, iterator)
}

//大于等于某值的集合
func (t *BTree) AscendGreaterOrEqual(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, pivot, nil, true, false, iterator)
}

func (t *BTree) Ascend(iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, nil, nil, false, false, iterator)
}

//降序方法
//大于等于，小于范围内集合
func (t *BTree) DescendRange(lessOrEqual, greaterThan Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, lessOrEqual, greaterThan, true, false, iterator)
}

//小于某值的集合
func (t *BTree) DescendGreaterThan(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, nil, pivot, false, false, iterator)
}

//大于等于某值的集合
func (t *BTree) DescendLessOrEqual(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, pivot, nil, true, false, iterator)
}

func (t *BTree) Descend(iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, nil, nil, false, false, iterator)
}

//获取key值
func (t *BTree) Get(key Item) Item {
	if t.root == nil {
		return nil
	}
	return t.root.get(key)
}

//获取最小值
func (t *BTree) Min() Item {
	return min(t.root)
}

//获取最大值
func (t *BTree) Max() Item {
	return max(t.root)
}

//判断是否有该值
func (t *BTree) Has(key Item) bool {
	return t.Get(key) != nil
}

//获取长度
func (t *BTree) Len() int {
	return t.length
}

func (t *BTree) Clear(addNodesToFreelist bool) {
	if t.root != nil && addNodesToFreelist {
		t.root.reset(t.cow)
	}
	t.root, t.length = nil, 0
}
