// goroutine 执行顺序和同步
// GMP G=goroutine P=process 管理n个goroutine队列，减少堵塞 M=go对线程的虚拟的实现
//P个数默认是cpu默认线程数
//goroutine调度是在用户态,不要在用户态和内核态频繁切换，轻量级的线程
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	// defer wg.Add(-1)
	fmt.Println("hello goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait() //等待
	//输出顺序不一定
}
