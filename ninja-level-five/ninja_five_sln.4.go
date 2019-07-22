package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	s := struct {
		vehicle
		xover bool
	}{
		vehicle: vehicle{
			doors: 4,
			color: `orange`,
		},
		xover: true,
	}

	fmt.Println("New Car: ", s)
}
