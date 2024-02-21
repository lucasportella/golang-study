package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("Recovered from the following panic: %v", r)
		}

	}()
	panicAbove3(0)
}

func panicAbove3(value int) {
	if value > 3 {
		panic("Value above 3, panicking!")
	}
	panicAbove3(value + 1)
}
