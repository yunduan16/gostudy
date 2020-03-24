//指针
//new 会返回类型的指针，并初始化类型的零值
//var a *int 只会声明，但是不初始化
//new 格式：func new(Type) *Type
package main

import "fmt"

func main() {
	var a *int
	fmt.Printf("Type=%T,value=%v\n", a, a) //value=nil
	a = new(int)
	*a = 100000000000
	b := &a
	fmt.Println(*a, a, &a) //a是指针类型， a输出的地址, &a取出a的地址的地址，a和&a不一样
	fmt.Println(*b)

	var c = 12
	// fmt.Println(c, *c) //错误，只有地址类型才可以用*，根据地址取值
	fmt.Println(c, &c) //错误，只有地址类型才可以用*，根据地址取值
}
