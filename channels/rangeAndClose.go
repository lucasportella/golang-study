package main

import "fmt"

func main() {
    // Create an integer channel
    ch := make(chan int)

    // Start a goroutine to send values into the channel
    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i // Send value into the channel
        }
        close(ch) // Close the channel after sending all values
    }()

    // Receive values from the channel in the main goroutine
    for num := range ch {
        fmt.Println(num) // Print received value
    }
}