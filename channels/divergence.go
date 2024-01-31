package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	go send(canal1)
	go other(canal1, canal2)
	for v := range canal2 {
		fmt.Println(v)
	}
}

func send(canal1 chan int) {
	for i := 1; i <= 10; i++ {
		canal1 <- i
	}
	close(canal1)
}

func other(canal1 chan int, canal2 chan int) {
	for v := range canal1 {
		wg.Add(1)
		go func(v int) {
			//channel1 data is split into several go routines and then all these data will converge to the channel 2 and be printed
			canal2 <- trabalho(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(v int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return v
}
