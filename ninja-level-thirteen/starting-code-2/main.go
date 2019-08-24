package main

import (
	"fmt"

	"github.com/genslein/go-programming/ninja-level-thirteen/starting-code-2/quote"
	"github.com/genslein/go-programming/ninja-level-thirteen/starting-code-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
