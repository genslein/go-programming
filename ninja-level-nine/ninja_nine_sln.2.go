package main

import (
	"fmt"
	"sync"
)

type person struct {
	First string
	Last  string
	Age   int64
}

func (p *person) speak() {
	fmt.Println("My name is", p.First, p.Last, "and I am", p.Age, "years old.")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("First goroutine, woohoo!")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second goroutine, yeehaw!")
		wg.Done()
	}()

	fmt.Println("Hello Playground")

	p := person{
		"James",
		"Bond",
		35,
	}
	saySomething(&p)
	// saySomething(p)
	wg.Wait()
}
