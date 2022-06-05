// 这个示例程序展示如何使用 atomic 包里的
// Store 和 Load 类函数来提供对数值类型
// 的安全访问
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown    int64
	waitGroup16 sync.WaitGroup
)

func main15() {
	waitGroup16.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	waitGroup16.Wait()
}

func doWork(name string) {
	defer waitGroup16.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
