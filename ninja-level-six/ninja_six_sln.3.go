package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Deferred message: I finished Printing!")
	fmt.Println(furlSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
	fmt.Println(arrSum([]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}))
}

func furlSum(y ...int) int {
	x := 0
	for _, v := range y {
		x += v
	}
	return x
}

func arrSum(y []int) int {
	x := 0
	for _, v := range y {
		x += v
	}
	return x
}
