package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2.0)
}

func (c circle) area() float64 {
	return (math.Pi * math.Pow(c.radius, 2))
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("my area is:", s.area())
}

func main() {
	x := square{
		length: 6.0,
	}
	y := circle{
		radius: 3.0,
	}

	info(x)
	info(y)
	return
}
