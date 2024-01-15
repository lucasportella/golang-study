package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	sliceEven := []int{}
	sliceOdd := []int{}
	var wg sync.WaitGroup
	go send(even, odd, quit, &wg)
	wg.Add(1)
	receive(even, odd, quit, &sliceEven, &sliceOdd)
	wg.Wait()
	close(quit)
}

func send(even chan int, odd chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- true
}

func receive(even chan int, odd chan int, quit chan bool, sliceEven *[]int, sliceOdd *[]int) {
	for {
		select {
		case number := <-even:
			*sliceEven = append(*sliceEven, number)
		case number := <-odd:
			*sliceOdd = append(*sliceOdd, number)
		case <-quit:
			fmt.Printf("Even: %v\n", *sliceEven)
			fmt.Printf("Odd: %v\n", *sliceOdd)
			return
		}
	}
}
