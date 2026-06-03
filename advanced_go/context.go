//go:build ignore

package main

import (
	"context"
	"fmt"
)

// 1. Create a custom type for your keys to prevent collisions!
// (The instructor noted that using basic strings like "userID" is dangerous
// because someone else might use the same string and overwrite your data)
type contextKey string

const userIDKey contextKey = "userID"

// 2. The downstream function that needs the data
func fetchTrucks(ctx context.Context) {
	// Extract the value using ctx.Value() and type assert it
	userID := ctx.Value(userIDKey).(int)
	fmt.Printf("Fetching trucks for User ID: %d\n", userID)
}

func main() {
	// 3. context.Background() creates an empty, top-level context
	// This is the starting point for 99% of Go programs.
	bgCtx := context.Background()

	// 4. context.WithValue creates a NEW context containing our data
	ctxWithUser := context.WithValue(bgCtx, userIDKey, 42)

	// Pass the new context down the chain
	fetchTrucks(ctxWithUser)
}
