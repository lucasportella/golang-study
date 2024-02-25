// Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// - Tire estes números do canal e demonstre-os.

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	chan1 := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				chan1 <- i
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(chan1)
	}()

	for val := range chan1 {
		fmt.Println(val)
	}
}
