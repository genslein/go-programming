package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	const gs = 100
	counter = 0
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			runtime.Gosched()
			atomic.AddInt64(&counter, int64(1))
			fmt.Println("New counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
