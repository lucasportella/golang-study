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
	go func1()
	func2()
}

// outputs be like:
// 115685
// -------------------
// 115686
// -------------------
// 115687
// 115688
// -------------------
// 115689
// -------------------
// 115690
// -------------------
// -------------------
// -------------------
// -------------------
// 115691
// 115692
// 115693
// -------------------
// 115694
// -------------------
// 115695
// 115696
// -------------------
