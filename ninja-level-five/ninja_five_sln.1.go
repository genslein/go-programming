package main

import (
	"fmt"
)

type person struct {
	first, last     string
	iceCreamFlavors []string
}

func main() {
	var ppl = []person{
		{
			first:           "joey",
			last:            "chesnut",
			iceCreamFlavors: []string{"hot dogs", "hamburgers"},
		},
		{
			first:           "kobayashi",
			last:            "maroon",
			iceCreamFlavors: []string{"tobiko", "green tea"},
		},
	}
	for _, v := range ppl {
		fmt.Println(v.first, v.last, `Favorite Ice Cream Flavors: `)
		for _, x := range v.iceCreamFlavors {
			fmt.Println("\t", x)
		}
	}
}
