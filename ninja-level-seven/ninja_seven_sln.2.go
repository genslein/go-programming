package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	age     int
	address string
}

func changeMe(p *person) {
	fmt.Println("Value of person before:", (*p))
	p.address = "1 Sea Anemone Lane, Great Barrier Reef, Airlie, AU"
}

func main() {
	x := "Hello World"
	y := person{
		"Dory",
		"Bluefin",
		5,
		"42 Wallaby Way, Sydney, AU",
	}
	changeMe(&y)
	fmt.Println("address of x:", &x)
	fmt.Println("Value of person after:", y)
}
