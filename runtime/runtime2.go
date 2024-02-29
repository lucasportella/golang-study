package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Your system is %v and your architecture is %v.", runtime.GOOS, runtime.GOARCH)
}