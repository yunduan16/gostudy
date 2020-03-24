//函数相关
package main

import (
	"errors"
	"fmt"
	"strings"
)

func swap(x, y int) (int, int) {
	return y, x
}

//定义函数类型  type 类型名称 类型结构
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

//返回值中有函数类型
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//传参中有函数变量
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	fmt.Println(swap(1, 2))

	var c calculation
	c = add

	f := sub
	fmt.Printf("c type of:%#T\n", c) //calculation 类型
	fmt.Printf("f type of:%#T\n", f) //func类型
	fmt.Println(c(2, 5))

	rs, err := do("+")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs(12, 13))

	rs2 := calc(5, 2, sub)
	fmt.Println(rs2)

	var f2 = adder()
	// var f2 = adder	//错误，因为这样相当于函数体赋值，但是该函数没有参数
	fmt.Println(f2(10)) //10
	fmt.Println(f2(20)) //30
	fmt.Println(f2(30)) //60
	//在f2的生命周期内，变量x一直有效
	f2 = adder()
	fmt.Println(f2(-20)) //-20

	jpgFunc := makeSuffixFunc(".jpeg")
	pngFunc := makeSuffixFunc(".png")
	fmt.Println(jpgFunc("test1"))
	fmt.Println(pngFunc("test2"))
}
