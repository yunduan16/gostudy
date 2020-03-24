// defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
//所以，如果defer调用的函数参数未确定，会先执行，一旦确定不会再修改，因为是值传递
//因为defer会对函数体的指针和参数进行拷贝，入栈(内存)
//defer 在panic发生后会执行，os.Exit(0)除外
//调用os.Exit时defer不会被执
package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	/*
		输出：
		A
		1
		2
		3
		B
		10
		2
		12
		BB
		10
		12
		22
		AA
		1	//这里不是10，因为这个函数参数已经传入 calc("AA", x, calc("A", x, y)) ,此时x=1，不会变化
		3
		4
	*/
}
