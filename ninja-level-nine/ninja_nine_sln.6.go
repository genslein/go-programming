package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go running on architecture:", runtime.GOARCH)
	fmt.Println("Go running on OS:", runtime.GOOS)
}
