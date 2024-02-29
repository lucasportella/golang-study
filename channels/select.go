package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start...")
	canal1 := make(chan int)
	canal2 := make(chan int)
	x := 10
	go func(x int) {
		defer close(canal1)
		for i := 0; i < x; i++ {
			canal1 <- i
		}
		fmt.Println("Closing channel1")
	}(x)

	go func(x int) {
		defer close(canal2)
		for i := 10; i < x; i++ {
			canal2 <- i
		}
		fmt.Println("Closing channel2")
	}(x * 2)

	for i := 0; i < x*3; i++ {
		select {
		case v := <-canal1:
			fmt.Println("Canal1: ", v)
		case valor := <-canal2:
			fmt.Println("Canal2: ", valor)
		}
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAA")

	}
}
