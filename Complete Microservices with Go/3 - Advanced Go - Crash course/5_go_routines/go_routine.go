//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

// processTruck simulates a slow task, like hitting a database or an API.
func processTruck(id int) {
	// Pauses this specific task for exactly 1 second
	time.Sleep(1 * time.Second)
	fmt.Printf("Truck %d processed.\n", id)
}

func main() {
	// A simple list of 5 trucks to process
	trucks := []int{1, 2, 3, 4, 5}

	// ==========================================================
	// 1. SYNCHRONOUS EXECUTION (The Slow Way)
	// ==========================================================
	fmt.Println("--- Starting Synchronous Process ---")

	// Start a timer
	startSync := time.Now()

	// Process one by one
	for _, truck := range trucks {
		processTruck(truck)
	}

	// Calculate how much time has passed
	syncDuration := time.Since(startSync)
	fmt.Printf("Synchronous Execution Time: %v\n\n", syncDuration)

	// ==========================================================
	// 2. CONCURRENT EXECUTION (The Fast Way with Goroutines)
	// ==========================================================
	fmt.Println("--- Starting Concurrent Process ---")

	// Start a new timer
	startConcurrent := time.Now()

	var wg sync.WaitGroup

	for _, truck := range trucks {
		wg.Add(1)

		go func(t int) {
			processTruck(t)
			wg.Done()
		}(truck)
	}

	// Wait for all 5 Goroutines to finish
	wg.Wait()

	// Calculate how much time has passed
	concurrentDuration := time.Since(startConcurrent)
	fmt.Printf("Concurrent Execution Time: %v\n", concurrentDuration)
}
