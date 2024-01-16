package main

import "fmt"

func main() {
	chan1 := make(chan int)
	go func() {
		// sempre temos que enviar um valor a um canal. Nesse caso estamos apenas enviando 1 único valor (0) e o canal depois fecha
		chan1 <- 0
		close(chan1)
	}()
	
	// recebemos o valor do canal1, se tiver valor inicializado, vai ser ok (é o caso)
	v, ok := <-chan1
	if ok {
		fmt.Printf("Comma ok: %v, value: %v\n", ok, v)
	} else {
		fmt.Printf("Comma not ok: %v, value: %v\n", ok, v)
	}

	// esperamos receber um segundo valor do canal1, mas ele só enviou 1 valor e depois fechou, logo o go manda o valor falsy padrão de um channel de inteiros, que é zero. Como não foi inicializado, o ok é falso.
	v, ok = <-chan1
	if ok {
		fmt.Printf("Comma ok: %v, value: %v\n", ok, v)
	} else {
		fmt.Printf("Comma not ok: %v, value: %v\n", ok, v)
	}
}
