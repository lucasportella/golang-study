package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	sliceEven := []int{}
	sliceOdd := []int{}
	go receive(even, odd, quit, &sliceEven, &sliceOdd)
	for i := 0; i <= 50; i++ {
		send(even, odd, quit, i)
	}
}

func send(even chan int, odd chan int, quit chan bool, number int) {
	if number == 50 {
		even <- number
		quit <- true
	} else if number%2 == 0 {
		even <- number
	} else {
		odd <- number
	}
}

func receive(even chan int, odd chan int, quit chan bool, sliceEven *[]int, sliceOdd *[]int) {
	for {
		select {
		case number := <-even:
			*sliceEven = append(*sliceEven, number)
		case number := <-odd:
			*sliceOdd = append(*sliceOdd, number)
		case <-quit:
			fmt.Printf("Even: %v\n", sliceEven)
			fmt.Printf("Odd: %v\n", sliceOdd)
			return
		}
	}
}
