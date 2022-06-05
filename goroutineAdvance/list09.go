package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter   int64
	WaitGroup sync.WaitGroup
)

func main09() {
	WaitGroup.Add(2)
	go incCounter(1)
	go incCounter(2)
	WaitGroup.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer WaitGroup.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
