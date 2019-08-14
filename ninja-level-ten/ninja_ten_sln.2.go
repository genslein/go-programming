package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cs := (chan<- int)(c)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	cr := (<-chan int)(c)

	go func() {
		c <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
