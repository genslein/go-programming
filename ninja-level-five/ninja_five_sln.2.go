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

	m := map[string]person{
		ppl[0].last: ppl[0],
		ppl[1].last: ppl[1],
	}

	for k, v := range m {
		fmt.Println(v.first, "key '"+(k)+"' matches "+v.last, `Favorite Ice Cream Flavors: `)
		for _, x := range v.iceCreamFlavors {
			fmt.Println("\t", x)
		}
	}
}
