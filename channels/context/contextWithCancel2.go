package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
							// gen(ctx) === dst chan int
	for n := range gen(ctx) {
		// the "for" above ranges through dst channel
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				// this case cannot happen because it is not sending a value to dst, and only dst is ready to receive values. Only when main function finishes, cancel() function will execute. The expression <-ctx.Done() will evaluate that the context is cancelled and this case statement should be triggered
				return // returning not to leak the goroutine
			case dst <- n:
				// since there is a channel ready to receive values, this case happens and "n" value is sent to the dst channel
				n++
			}
		}
	}()
	return dst
}


// "<-ctx.Done()"" acts as an if statement that checks if context has been cancelled or not