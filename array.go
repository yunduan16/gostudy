package main

import "fmt"

func main() {
	//数组长度一旦确定不能增删
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "广州", "深圳"}
	arr := [...]int{1: 1, 3, 5: 6, 8}

	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)
	fmt.Println(arr)

	fmt.Println("------------数组arr的遍历------------")
	for index, value := range arr {
		// for i := 0;i < len(a);i++ {
		fmt.Println(index, value)
	}
	fmt.Println("------------二维数组的遍历------------")
	//多维数组只能在第一层才可用...表示  array[...][3]
	arr2 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "杭州"},
	}

	fmt.Println(arr2[1])    //广州 深圳
	fmt.Println(arr2[2][1]) //杭州
}
