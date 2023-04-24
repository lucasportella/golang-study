package main

import (
	"fmt"
	"time"
)

func func1() {
	for i := 0; i < 500000; i++ {
		fmt.Println(i)
		time.Sleep(1000)
	}
}

func func2() {
	for i := 0; i < 500000; i++ {
		fmt.Println(i)
		time.Sleep(1000)
	}
}

func main() {
		func1()
		func2()
}

