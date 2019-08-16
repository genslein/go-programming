package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("This is my error number: %d", ce.info)
}

func foo(c error) {
	fmt.Println("foo ran -", c, "\n")
}

func main() {
	fmt.Println("hello world")
	e := customErr{
		info: "need more coffee",
	}

	foo(e)
}
