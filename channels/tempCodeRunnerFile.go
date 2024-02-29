package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	trabalho()
	wg.Wait()
}

func trabalho() chan string {
	chan1 := make(chan string)
	go func() {
		defer wg.Done() // Ensure that Done() is called when the goroutine exits
		for i := 1; i <= 100; i++ {
			chan1 <- fmt.Sprintf("loop value to %v: %v", chan1, i)
			time.Sleep(time.Duration(rand.Intn(1e3)))
		}
	}()
	return chan1
}
