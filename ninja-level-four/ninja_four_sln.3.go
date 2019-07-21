package main

import (
	"fmt"
)

func main() {
	var c = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(c[:5])
	fmt.Println(c[5:10])
	fmt.Println(c[2:7])
	fmt.Println(c[1:6])

	fmt.Printf("%T\n", c)
}
