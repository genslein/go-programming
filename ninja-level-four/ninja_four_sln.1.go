package main

import (
	"fmt"
)

func main() {
	var c = []int{1, 2, 3, 4, 5}

	for i, v := range c {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", c)
}
