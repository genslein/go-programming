package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func(j int) {
			for i := 0; i < 10; i++ {
				c <- (j * 10) + i
			}
		}(j)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("Done. Exiting")
}
