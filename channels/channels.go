package main

import (
	"fmt"
)

func send(ch chan<- int) {
	ch <- 40 
}

func receive(ch <-chan int) {
	fmt.Printf("valor recebido de um canal: %v", <-ch)
}

func main() {
	canal1 := make(chan int)
	go send(canal1)
	receive(canal1)
}
