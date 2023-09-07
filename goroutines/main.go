package main

import (
	"fmt"
	"time"
)

func DoAddition(a, b int, ch chan int) {
	fmt.Println("I'm doing addition")
	time.Sleep(2 * time.Second)
	result := a + b
	ch <- result
	fmt.Println("Result is", result)
}

func main() {
	ch := make(chan int)

	// addition 1
	go DoAddition(1, 2, ch)
	// addition 2
	go DoAddition(3, 4, ch)
	// addition 3
	go DoAddition(5, 6, ch)

	// obtain the results
	r1 := <-ch		// temporal deadlock
	r2 := <-ch		// temporal deadlock
	r3 := <-ch		// temporal deadlock

	r4 := <-ch		// permanent deadlock

	fmt.Println("Results are", r1, r2, r3, r4)

	fmt.Println("Program has finished executing")
}