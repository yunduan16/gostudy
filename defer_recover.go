//recover函数必须配合defer使用，defer定义在前
package main

import "fmt"

func f1() {
	fmt.Println("func f1")
}

func f2() {
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("recover in f2")
	// 	}
	// }()
	panic("func f2 exit") //不会执行
	/*
		//recover语句在panic之后不起作用
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("recover in f2")
			}
		}()
	*/
}

func f3() {
	fmt.Println("func f3")
}
func main() {
	f1()
	f2()
	f3()
}
