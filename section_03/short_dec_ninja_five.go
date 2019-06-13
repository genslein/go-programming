package main

import (
	"fmt"
)

type hotdog int

var x hotdog
var y int
var z = 43

func main() {
	fmt.Printf("%v\n%T\n", x, x)
	x = 42
	fmt.Printf("%v\n%T\n", x, x)
	y = int(x)
	fmt.Printf("%v\n%T\n", y, y)
	fmt.Printf("%v\n%T\n", z, z)
}
