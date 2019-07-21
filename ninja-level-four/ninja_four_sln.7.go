package main

import (
	"fmt"
)

func main() {
	var x = [][]string{{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, v := range x {
		fmt.Println(i, v)

		for k, w := range v {
			fmt.Println(i, k, w)
		}
	}
}
