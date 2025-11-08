package main

import (
	"fmt"
	"sync"
	"time"
)

// task simulates a job that takes some time to complete.
// It takes a name for identification and a pointer to WaitGroup
func task(name string, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure Done() is called when function exits

	for i := 1; i <= 3; i++ {
		// Print the iteration
		fmt.Println(name, "iteration", i)
		// Simulate work with sleep
		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	// Create a WaitGroup
	var wg sync.WaitGroup

	// We are going to start 2 goroutines, so set counter = 2
	wg.Add(2)

	// Launch first goroutine
	go task("A", &wg)

	// Launch second goroutine
	go task("B", &wg)

	// Wait for all goroutines to finish
	// main will pause here until counter becomes 0
	wg.Wait()

	// All goroutines finished
	fmt.Println("All done!")
}
