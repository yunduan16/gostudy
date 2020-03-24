//互斥锁
//sync除了Once,WaitGroup，其他都不建议使用，推荐用channel
//如果不加互斥锁 sync.Lock 则输出结果不一定
package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup //不能拷贝，独一份
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2) //开启两个进程
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
