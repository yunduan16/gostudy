//包含defer语句
package main

import (
	"fmt"
)

//deferDemo1 defer语句先进后出形式
func deferDemo1() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
		// } //去掉(x)，这里就只是函数体了
	}(x)
	//这里表示匿名函数
	return 5
}

//return 分为赋值给函数体返回值，RET退出函数操作，defer执行在两者之间
func main() {
	// deferDemo1()
	//函数体内部的defer作用域只在函数内，不在函数外起作用
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	y := 1
	func(y int) {
		y++
	}(y)
	fmt.Println("y=", y)
}
