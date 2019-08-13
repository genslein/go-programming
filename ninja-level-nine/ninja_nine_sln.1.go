package main

import (
	"fmt"
)

func main() {

	go func() {
		fmt.Println("First goroutine, woohoo!")
	}()

	go func() {
		fmt.Println("Second goroutine, yeehaw!")
	}()

	fmt.Println("Hello Playground")
}
