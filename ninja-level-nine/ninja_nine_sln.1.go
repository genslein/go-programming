package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("First goroutine, woohoo!")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second goroutine, yeehaw!")
		wg.Done()
	}()

	fmt.Println("Hello Playground")
}
