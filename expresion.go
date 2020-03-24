package main

import (
	"fmt"
	"time"
)

//switch 可以用表达式
func switchDemo() {
	a := 30
	//如果switch case用表达式,switch后面可以不用变量
	switch {
	case a < 18:
		fmt.Println("你是少年")
	case a >= 18 && a < 50:
		fmt.Println("你是青年")
	case a >= 50:
		fmt.Println("你已经步入老年")
	default:
		fmt.Println("活着真好")
	}
}

//兼容c语言的多个case公用一个break 使用fallthrough

func switchDemo2() {
	s := "c"
	switch {
	case s == "a": //不支持全等于 ===
		fmt.Println("a")
		fallthrough //从上往下顺序
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fallthrough
	case s == "d":
		fmt.Println("c and d")
	default:
		fmt.Println("none")
	}
}

func ifDemo1() {
	// score := 10
	if score := 65; score >= 90 {
		fmt.Println("score is A")
	} else if score > 75 {
		fmt.Println("score is B")
	} else {
		fmt.Println("score is C")
	}
	// fmt.Printf("score=%v \n", score) //undefined 因为score是只有在if语句中才生效
}

//for循环省略初始和结束语句，前后分号;都不需要加
func forDemo1() {
	i := 0
	for i < 20 {
		fmt.Println("i=", i)
		i++
	}
}

//无限循环
func forDemo2() {
	for {
		fmt.Println("有个进程一直在等待你开启，5秒开启一次")
		// time.Sleep(time.Microsecond * 100)
		time.Sleep(time.Second * 5)
	}
}

//break语句指定跳出循环的代码块
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

//select 语句case必须是一次io操作(读写)
// 与switch语句相比， select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，大致的结构如下
/*
	select {
	case <-chan1:
		// 如果chan1成功读到数据，则进行该case处理语句
	case chan2 <- 1:
		// 如果成功向chan2写入数据，则进行该case处理语句
	default:
		// 如果上面都没有成功，则进入default处理流程
	}
*/
func main() {
	switchDemo()

	switchDemo2()

	ifDemo1()

	forDemo1()

	breakDemo1()
	// forDemo2()
}
