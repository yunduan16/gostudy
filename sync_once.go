//sync.Once只执行一次
// Once.Do(func)
//sync.Map支持并发安全写入
//sync.atomoic基本类型的原子性操作方法，了解
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var icons map[string]string
var loadIconsOnce sync.Once
var m sync.Map

func loadIcons()  {
	icons = map[string]string{
		"left": "left.png",
		"up": "up.png",
		"right": "right.png",
		"down": "down.png",
	}
}

func Icon(name string) string  {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func get(key string) (interface{}, bool) {
	return m.Load(key)
}

func set(key string, value int)  {
	m.Store(key, value)
}
func syncMap() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func (n int)  {
			key := strconv.Itoa(n)	//转换为10进制的字符串表示
			// fmt.Printf("%T\n", key)
			set(key, n)
			value, _ := get(key)
			fmt.Printf("k=%v,v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main()  {
	// name := "down"
	// fmt.Println(Icon(name))

	syncMap()
}