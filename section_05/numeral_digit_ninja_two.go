package main

import (
	"fmt"
)

var x int = 42

func main() {
	a := x == 42
	b := x <= 42
	c := x != 42
	d := x < 42
	e := x > 42
	f := x >= 42

	fmt.Printf("%t\t%t\t%t\t%t\t%t\t%t\n", a, b, c, d, e, f)
}
