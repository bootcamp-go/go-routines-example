package main

import (
	"fmt"
	"time"
)

func DoAddition(a, b int) {
	fmt.Println("I'm doing addition")
	time.Sleep(2 * time.Second)
	result := a + b
	fmt.Println("Result is", result)
}

func main() {
	// addition 1
	go DoAddition(1, 2)
	// addition 2
	go DoAddition(3, 4)
	// addition 3
	go DoAddition(5, 6)

	fmt.Println("Program has finished executing")
}