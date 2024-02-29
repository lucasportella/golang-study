package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int { // "q" channel direction not changed
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int, q chan int) {
	for val := range c { // uses direct loop to show values and end function. Since for loop without select knows when channel is closed, the receive function doesn't need to consume to quit channel
		fmt.Println(val)
	}
}