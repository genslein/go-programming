package main

import (
	"fmt"
)

func main() {
	var c = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	c = append(c, 52)
	fmt.Println(c)
	c = append(c, 53, 54, 55)
	fmt.Println(c)
	y := []int{56, 57, 58, 59, 60}
	c = append(c, y...)
	fmt.Println(c)
}
