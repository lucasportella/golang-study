package main

import (
	"fmt"
)

func func1() {
	for i := 0; i < 500000; i++ {
		fmt.Println(i)
	}
}

func func2() {
	for i := 0; i < 500000; i++ {
		fmt.Println("-------------------")
	}
}

func main() {
		func1()
		func2()
}

