package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go") // create context With value from parent context context.Background()
	// the key of the value is "languague" and the value is "Go"

	foo(ctx, "language") // key "language" found with value "Go"

	ctx = context.WithValue(ctx, "dog", "Gaston") // adds one more key/value to the context

	foo(ctx, "dog") // key "dog" found with value "Gaston"
	foo(ctx, "language") // key "language" with its value "Go" is still in this context
	foo(ctx, "color") // key "color" not found
}

func foo(ctx context.Context, key string) {
	if value := ctx.Value(key); value != nil { // receive value from context through the key and check if is nil (short statement if)
		fmt.Println("found value:", value) // case key is found, shows its value
		return
	}
	fmt.Println("key not found:", key) // case key is not found, show unknown key
}
