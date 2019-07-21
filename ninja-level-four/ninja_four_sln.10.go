package main

import (
	"fmt"
)

func main() {
	var x = map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["goldfinger"] = []string{`The most beautiful death`, `Gold`, `Lasers`}

	delete(x, `moneypenny_miss`)
	for i, v := range x {
		fmt.Println(i, v)

		for k, w := range v {
			fmt.Println(i, k, w)
		}
	}
}
