package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	even := make(chan int)
	odd := make(chan int)
	converge := make(chan int)
	wg.Add(3)
	go receiveEven(even, converge)
	go receiveOdd(odd, converge)
	go send(even, odd)
	go func() {
		wg.Wait()       // must wait only to close the converge channel only, any synchronous wait will generate deadlock (convergence channel would be closed before receiveEven and Odd could send values to there)
		close(converge) // need to close converge in a go func or it wont be able to read from channel

	}()

	for v := range converge {
		fmt.Println(v)
	}
}

func send(even chan int, odd chan int) {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	wg.Done()
}

func receiveEven(even chan int, converge chan int) {
	for v := range even {
		converge <- v
	}
	wg.Done()
}
func receiveOdd(odd chan int, converge chan int) {
	for v := range odd {
		converge <- v
	}
	wg.Done()
}
