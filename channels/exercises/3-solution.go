package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int) // its a bidirectional channel, but this function will return this channel as a receive only channel. That's why the goroutine below can send value through the channel, because it uses the "c" declaration of bidirectional channel
	go func() { // go routine so the gen fn can return c and be used by receive fn
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // close channel to sign this channel will not send anymore values
	}()
	return c
}

func receive(canal <-chan int) {
	for v := range canal { // loop through channel so we can receive its values and print them
		fmt.Println(v)
	}
}
