package main

import (
	"fmt"
	"runtime"
	"time"
)

func a()  {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b()  {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func c()  {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func main()  {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	go c()
	time.Sleep(time.Second)
}