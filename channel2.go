package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	//在主goroutine中获取ch2的数据
	//单向的读写都不影响for range的获取数据
	for i := range ch2 { //ch2关闭会退出for循环
		fmt.Println(i)
	}
}
