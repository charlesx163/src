// 这个示例程序展示如何使用互斥锁来
// 定义一段需要同步访问的代码临界区
// 资源的同步访问
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counterm int
	wgm      sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	wgm.Add(2)
	go incCounterMutex(1)
	go incCounterMutex(2)
	wgm.Wait()
	fmt.Printf("Final Counter: %d\n", counterm)
}

func incCounterMutex(id int) {
	defer wgm.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counterm
			runtime.Gosched()
			value++
			counterm = value
		}
		mutex.Unlock()
	}
}
