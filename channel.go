//go提倡通过通信共享内存
//ch := make(chan int)无缓存通道，又称为同步通道
package main

import "fmt"

func recv(c chan int)  {
	ret := <- c
	fmt.Println("接收成功：", ret)
}

func main()  {
	var ch chan int//未进行初始化，需要make做初始化，只声明的通道不能接收数据
	fmt.Printf("%#T, %#v %p\n", ch, ch, ch)
	ch = make(chan int)	//make会对通道进行初始化
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
	/*
	ch <- 10	//发送数据到通道
	v := <- ch	//从通道接收数据
	<-ch		//无接受者接收通道数据
	close(ch) //关闭通道

	*/
}