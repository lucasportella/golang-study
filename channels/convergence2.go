/* 
- Func trabalho cria um canal, cria uma go func que manda dados pra esse canal, e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
- Func converge toma dois canais, cria um canal novo, e cria duas go funcs com for infinito que passa tudo para o canal novo. Retorna o canal novo.
- Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	converged := converge(work(), work())
	for i := 1; i <= 200; i++ {
		fmt.Printf("Value from converged: %v\n", <-converged)
	}
}

func work() chan string {
	newChan := make(chan string)
	go func() {
		for i := 1; i <= 100; i++ {
			newChan <- fmt.Sprintf("loop value: %v\n", i)
			time.Sleep(time.Duration(rand.Intn(1e3)))
		}
	}()
	return newChan
}

func converge(chan1 chan string, chan2 chan string) chan string {
	converge := make(chan string)
	go func() {
		for {
			value := <-chan1
			converge <- value
		}
	}()
	go func() {
		for {
			// <- <- means "receive value from channel and immediately send it to another channel", we can't just use one "<-" because converge channel must receive a string value, not a chan value
			converge <- <- chan2
		}
	}()
	return converge
}
