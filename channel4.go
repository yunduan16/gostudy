//用通道实现简单的work pool
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启jobs的goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)
	//输出结果
	for k := 1; k <= 5; k++ {
		<-results
	}
}
