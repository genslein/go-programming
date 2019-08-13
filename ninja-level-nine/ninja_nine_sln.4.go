package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	const gs = 100
	counter := 0
	var wg sync.WaitGroup
	var mux sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			runtime.Gosched()
			mux.Lock()
			counter++
			fmt.Println("New counter:", counter)
			mux.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
