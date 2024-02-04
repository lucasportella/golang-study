package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)                   // this sets a deadline of the current time plus 50 milliseconds
	ctx, cancel := context.WithDeadline(context.Background(), d) // sets the context with the deadline above

	defer cancel() // executed when main fn returns

	select { // gets locked till one case matches
	case <-time.After(1 * time.Second): // executed if the timeout of 1 second is reached before the context deadline, since 50 milli < 1 sec, it will not be executed
		fmt.Println("overslept")
	case <-ctx.Done(): // executed when context is cancelled, when the deadline of the context is reached the context will be cancelled(after the 50 milliseconds)
		fmt.Println(ctx.Err())
	}

}
