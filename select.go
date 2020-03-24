package main

import (
    "fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	fmt.Println("执行一次")
	// v := <-c
	// fmt.Printf("c=%v \n", v)
	// c <- v
    for {
		fmt.Println("死循环执行。。。")
        select {
		case c <- x:	//表示数据x 进入到管道，即数据的写过程，管道类似一个共享的文件
		// case v := <-c: //表示从管道c接收数据并赋值给v，<-c表示数据从管道出来，即数据的读过程
			x, y = y, x+y
			fmt.Printf("x=%v, y=%v \n", x, y)
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
	quit := make(chan int)
	//无缓存通道必须有一个读协程才可进行写入，先执行读协程，再执行写协程

    go func() {
        for i := 0; i < 8; i++ {
			fmt.Printf("第%v层循环 \n", i+1)
			fmt.Println(<-c)
        }
        quit <- 0	//当循环结束，quit有写入了，通道完整了，才会进行退出死循环操作
    }()

	fibonacci(c, quit)
}