package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())              // error check 1: <nil> (returns nil because context not yet closed)
	fmt.Println("num gortins 1:", runtime.NumGoroutine()) // num gortins 1: 1 (only main fn running as gortins)

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done(): // when context is cancelled this triggers when the go routine returns
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200) // 200 millisecond for each iterator of this goroutine
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2) // 2 seconds, the goroutine should have printed around 10 workings at this time
	fmt.Println("error check 2:", ctx.Err()) // error check 2: <nil> (still nil, context not yet cancelled)
	fmt.Println("num gortins 2:", runtime.NumGoroutine()) // num gortins 2: 2  (the gortin from line 16 and main fn)

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err()) // error check 3: context canceled
	fmt.Println("num gortins 3:", runtime.NumGoroutine()) // num gortins 3: 1 (since context is cancelled the gortin from line 16 returned)
}
