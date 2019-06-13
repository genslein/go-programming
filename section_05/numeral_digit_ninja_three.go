package main

import (
	"fmt"
)

const (
	x     = 42
	y int = 43
)

func main() {
	fmt.Printf("%v %T\t%v %T\n", x, x, y, y)
}
