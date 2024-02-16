// - Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			chan1 <- i
		}
		close(chan1)
	}()
	for val := range chan1 {
		fmt.Println(val)
	}
}
