package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
		fmt.Printf("a:%v type:%T len:%d cap:%d \n", a, a, len(a), cap(a))

		//切片会保留初始数组起始位置index=1后的数据
		//切片x:n(n不能大于原数组的长度，左包右开)
		b := a[1:3]
		fmt.Printf("b:%v type:%T len:%d cap:%d \n", b, b, len(b), cap(b))

		//切片再切片，b保留了a数组index=1后面所有的数据{"上海", "广州", "深圳", "成都", "重庆"}，
		//并将它存入新的底层数据中，所以c的index=1 => "广州"
		c := b[1:5]
		fmt.Printf("c:%v type:%T len:%d cap:%d \n", c, c, len(c), cap(c))

		//1,切片直接不能比较,切片!=nil并不代表为空，一般用len(slice) == 0判断为空
		var s1 []int
		s2 := []int{}
		s3 := make([]int, 0)

		fmt.Println(s1 == nil, s2 == nil, s3 == nil) //true, false, false
	*/
	sliceDemo1()

	sliceDemo2()

	sliceCopyDemo()

	sliceRemoveDemo()

	sliceTest1()

	sliceTest2()
}

//2,切片的赋值拷贝，会共用一套底层数组
func sliceDemo1() {
	s1 := make([]int, 3) //会初始化
	s2 := s1
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}

//切片的append扩容
func sliceDemo2() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		//cap发生修改后，地址发生修改
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

//切片的copy，1直接赋值(引用) 2copy函数（复制一份）
func sliceCopyDemo() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)

	copy(c, a)
	fmt.Println(a)
	fmt.Println(c)

	c[0] = 1000
	fmt.Println(a)
	fmt.Println(c)
}

//切片删除，用append和[n:m]形式
func sliceRemoveDemo() {
	a := []int{30, 31, 32, 33, 34, 35, 36}
	//要删除a中索引为2
	a = append(a[:2], a[3:]...)
	fmt.Printf("a=%v len=%d cap=%d\n", a, len(a), cap(a))
}

func sliceTest1() {
	var a = make([]string, 5, 10) //初始化默认为空
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a, len(a))
	fmt.Println(cap(a)) //20，在i=5时，需要扩容，翻倍10*2
}

//用sort包对数组进行排序
func sliceTest2() {
	var a = [...]int{3, 7, 8, 9, 1}
	//升序排序
	sort.Ints(a[:])
	fmt.Println(a)

	// IntSlice 格式化为数据接口
	// Reverse只是返回递减接口，并不排序
	sort.Sort(sort.Reverse(sort.IntSlice(a[:]))) //reverse只提供逆转接口和重写了less方法，swap方法
	fmt.Println(a)
}
