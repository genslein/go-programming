package main

import (
	"fmt"
)

func main() {
	i := 1988

	for {
		fmt.Println(i)
		if i++; i > 2019 {
			break
		}
	}
}
