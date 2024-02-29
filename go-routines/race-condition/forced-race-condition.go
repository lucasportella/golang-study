package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var contador int32
var wg sync.WaitGroup


func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(contador)
}
