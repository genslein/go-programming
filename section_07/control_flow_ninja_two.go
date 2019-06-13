package main

import (
	"fmt"
)

const (
	k = 64
)

func main() {
	for i := 1; i <= 26; i++ {
		fmt.Println(k + i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", k+i)
		}
	}
}
