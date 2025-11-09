package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function — runs as a goroutine
func worker(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Marks this worker as done when it finishes

	for i := 1; i <= 3; i++ {
		// Simulate some work with time.Sleep
		fmt.Printf("Worker %d doing iteration %d\n", id, i)
		time.Sleep(time.Duration(300+id*100) * time.Millisecond)
	}

	// Send message back to channel when finished
	ch <- fmt.Sprintf("✅ Worker %d finished all tasks", id)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 3) // Buffered channel (3 messages can fit)

	// Start 3 workers concurrently
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	// Goroutine to close channel once all workers done
	go func() {
		wg.Wait()   // wait until all 3 Done() called
		close(ch)   // close channel to signal no more messages
	}()

	// Continuously listen for messages from workers
	for {
		select {
		case msg, ok := <-ch:
			if !ok { // if channel closed
				fmt.Println("🎉 All workers finished!")
				return
			}
			fmt.Println("Received:", msg)
		default:
			// Non-blocking select — optional, shows progress
			time.Sleep(100 * time.Millisecond)
		}
	}
}
