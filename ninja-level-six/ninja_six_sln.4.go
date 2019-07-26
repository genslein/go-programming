package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age, "years old")
}

func main() {
	x := person{
		first: "Tom",
		last:  "Hanks",
		age:   60,
	}
	x.speak()
	return
}
