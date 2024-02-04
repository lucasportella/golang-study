package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond) // creates context that will cancel after 50 milliseconds timer finishes
	defer cancel()                                                                // cancel when main fn returns

	select { // waits till one of the cases matches
	case <-time.After(1 * time.Second): // cannot match because 1sec > 50 millisec
		fmt.Println("overslept")
	case <-ctx.Done(): // after the 50 milliseconds ctx is cancelled and it will match this case
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}
