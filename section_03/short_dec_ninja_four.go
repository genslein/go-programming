package main

import (
	"fmt"
)

type hotdog int

var x hotdog

func main() {
	fmt.Printf("%v\n%T\n", x, x)
	x = 42
	fmt.Println(x)
}
