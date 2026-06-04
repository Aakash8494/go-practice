//go:build ignore

package main

import (
	"fmt"
	"time"
)

// Function 1: Prints numbers 1 to 5
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		// Pause for half a second to simulate work
		time.Sleep(500 * time.Millisecond)
	}
}

// Function 2: Prints letters A to E
func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("Letter: %c\n", i)
		// Pause for half a second to simulate work
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting...")

	// 1. Start printNumbers as a goroutine (runs in the background)
	// printNumbers()
	go printNumbers()

	// 2. Run printLetters normally in the main program
	printLetters()

	// 3. Wait a moment before exiting so the background task can finish
	time.Sleep(1 * time.Second)

	fmt.Println("Done!")
}
