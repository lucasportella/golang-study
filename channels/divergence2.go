package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go send(chan1)
	go other(chan1, chan2)
	for v := range chan2 {
		fmt.Println(v)
	}
}

func send(chan1 chan int) {
	for i := 1; i <= 100; i++ {
		chan1 <- i
	}
	close(chan1)
}

func other(chan1 chan int, chan2 chan int) {
	// 5 values will be displayed per second(work fn takes 1sec), if chan1 received values are fewer than the go funcs (like 3 values), chan2 will only receive 3 values and display them
	// the higher the loop variable is from the below for, the faster the values will be displayed, in the case above is 5 per second
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			// this go func will range through chan1 100 values to send them to chan2, because there is another 4 go funcs doing the same job, every go func will like loop to only around 25 values
			for v := range chan1 {
				chan2 <- work(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(chan2)
}

func work(v int) int {
	time.Sleep(time.Millisecond * time.Duration(1000))
	return v
}
