package main

import (
	"fmt"
)

var x int = 42

func main() {
	fmt.Printf("%b\t%d\t%#x\n", x, x, x)
}
