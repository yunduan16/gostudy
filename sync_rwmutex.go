//互斥锁之读写锁
//读锁不会影响其他的读，但是堵塞写入，写锁进行时，会阻塞其他的读写
//sync.RWMutex
package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func write() {
	rwlock.Lock() //读写锁的写锁方法
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	rwlock.Unlock() //读写锁的解锁写锁的方法
	wg.Done()
}

func read() {
	rwlock.RLock() //读锁
	fmt.Println("read x value=", x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() //解除读锁
	wg.Done()
}
func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println("x=", x) //无论读写再怎么设置都不影响最终写入结果
}
