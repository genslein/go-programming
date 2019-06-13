package main

import (
	"fmt"
)

func main() {
	s := `"Today is going to be a beautiful day", said Peter Rabbit.
Henry Tortoise became silent. He was not convinced.

	"I do not agree with your assessment given the thunderstorm above our heads",
Henry Tortoise tersely replied.`

	//c := []byte(s); c.length < 60 is what I wanted to run here
	if false {
		fmt.Println(s)
	} else {
		fmt.Println("Peter Rabbit sighed dejectedly and returned to his rabbit hole.")
	}
}
