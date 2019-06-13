package main

import (
	"fmt"
)

type favorite string

var m favorite = "favSport"

func main() {
	switch m {
	case "favSport":
		fmt.Println("My favorite sport is American Football.")
	default:
		fmt.Println("This is the default output switch.")
	}
}
