package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 7
}

func bar() (int, string) {
	return 3, "bar returned"
}
