//go:build ignore

package main

import (
	"context"
	"fmt"
	"time"
)

// Always use custom keys so 3rd party packages don't overwrite your data!
type myKey string

const userIDKey myKey = "userID"

// ==========================================
// 3. ROUTE SERVICE (The slow database)
// ==========================================
func RouteService(ctx context.Context) {
	fmt.Println("   -> [Route Service] Starting to fetch routes...")

	// Simulating a super slow database query that takes 5 seconds
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("   -> [Route Service] Success! Here are the routes.")
	case <-ctx.Done():
		// Boom! The 2-second context timeout hit before the 5s task finished.
		fmt.Println("   -> [Route Service] ERROR:", ctx.Err())
	}
}

// ==========================================
// 2. TRUCK SERVICE (The main endpoint)
// ==========================================
func TruckService(ctx context.Context) {
	// USE CASE 1: METADATA
	// We extract the UserID that the Middleware secretly attached.
	userID := ctx.Value(userIDKey).(string)
	fmt.Printf("-> [Truck Service] Processing request for User: %s\n", userID)

	// USE CASE 2: TIMEOUTS
	// We need to call the Route Service, but we won't wait forever.
	// We wrap the current context with a 2-second ticking bomb.
	fmt.Println("-> [Truck Service] Calling Route Service (Timeout: 2s)")
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // Always clean up!

	// Pass the ticking context down to the next service
	RouteService(ctxWithTimeout)
}

// ==========================================
// 1. AUTH MIDDLEWARE (The entry point)
// ==========================================
func AuthMiddleware(bgCtx context.Context) {
	fmt.Println("[Auth Middleware] Verifying user token...")

	// Fake logged-in user we got from the database
	loggedInUser := "Bro_Aakash_99"

	// USE CASE 1: METADATA
	// We wrap the blank context with the User ID data.
	ctxWithUser := context.WithValue(bgCtx, userIDKey, loggedInUser)

	// Pass the stuffed context to the Truck Service
	TruckService(ctxWithUser)
}

func main() {
	fmt.Println("=== USER MAKES A GET REQUEST ===")

	// Start with a totally empty box
	bgCtx := context.Background()

	// Hand the box to the very first layer
	AuthMiddleware(bgCtx)

	fmt.Println("=== REQUEST FINISHED ===")
}
