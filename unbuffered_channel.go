package main

import "fmt"

func main() {
    ch := make(chan int) // create a channel

    // sender goroutine
    go func() {
        ch <- 42 // send value to channel
    }()

    // receiver in main goroutine
    val := <-ch
    fmt.Println("Received:", val)
}
