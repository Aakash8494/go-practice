//go:build ignore

package main

import (
	"context"
	"fmt"
	"time"
)

func processSlowTask(ctx context.Context) {
	// Simulating a database call that takes 5 seconds...
	// (Note: The select {} syntax is how Go listens for channels,
	// which he mentioned will be covered in depth in the next lecture!)
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task finished successfully!")
	case <-ctx.Done():
		// If the context timeout hits before the 5 seconds is up, this triggers!
		fmt.Println("Error:", ctx.Err()) // Prints: "context deadline exceeded"
	}
}

func main() {
	bgCtx := context.Background()

	// 1. context.WithTimeout wraps the parent context and adds a ticking clock.
	// It returns a new context, AND a manual cancel() function.
	ctxWithTimeout, cancel := context.WithTimeout(bgCtx, 2*time.Second)

	// 2. defer cancel() ensures that when this function finishes,
	// all background resources tied to this context are cleaned up (prevents memory leaks!)
	defer cancel()

	fmt.Println("Starting process (Timeout set to 2s, Task takes 5s)...")

	// Pass the timeout context into our function
	processSlowTask(ctxWithTimeout)
}
