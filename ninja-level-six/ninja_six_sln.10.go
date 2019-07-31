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

func (s square) cubify() func() float64 {
	return func() float64 {
		return s.length * s.length * s.length
	}
}

func cylindricize(ar func() float64, h float64) float64 {
	return ar() * h
}

func main() {
	x := square{
		length: 6.0,
	}
	y := circle{
		radius: 3.0,
	}

	{
		z := "This is my closure string"
		fmt.Println(z)
	}
	fmt.Println("I cannnot print my closure string here")

	func() {
		fmt.Println("This is the area of a triangle:", (6.0 * 3.0 * 0.5))
	}()

	rectangle := func(w int, l int) int {
		return w * l
	}

	fmt.Println(x.area())
	fmt.Println(y.area())
	fmt.Println("Area of a rectangle:", rectangle(5, 7))
	fmt.Println("The volume of a cube:", x.cubify()())
	fmt.Println("The volume of a cylinder", cylindricize(y.area, 10.0))
	return
}
