package main

import (
	"fmt"
)

func main() {
	var c = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	c = append(c[:3], c[6:]...)
	fmt.Println(c)
}
