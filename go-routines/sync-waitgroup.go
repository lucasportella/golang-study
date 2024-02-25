package main

import (
	"fmt"
	"sync"
	"time"
)

func func1() {
	for i := 0; i < 500000; i++ {
		fmt.Println(i)
		time.Sleep(1000)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 500000; i++ {
		fmt.Println("---------------")
		time.Sleep(1000)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func1()
	func2()
	wg.Wait()
}
