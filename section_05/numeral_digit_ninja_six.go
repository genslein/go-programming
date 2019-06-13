package main

import (
	"fmt"
)

const (
	a = 2016 + iota
	b = 2016 + iota
	c = 2016 + iota
	d = 2016 + iota
)

func main() {
	fmt.Printf("%v\t%v\t%v\t%v\n", a, b, c, d)
}
