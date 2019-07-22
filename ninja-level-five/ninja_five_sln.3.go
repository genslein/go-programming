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
	var t = truck{
		vehicle: vehicle{
			doors: 2,
			color: `black`,
		},
		fourWheel: true,
	}

	var s = sedan{
		vehicle: vehicle{
			doors: 4,
			color: `silver`,
		},
		luxury: false,
	}

	fmt.Println("Four wheel capable: ", t.fourWheel, t.color, t.doors)
	fmt.Println("Luxury sedan: ", s.luxury, s.color, s.doors)
}
