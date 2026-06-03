//go:build ignore

package main

import (
	"context"
	"fmt"
)

// The SAFE way: Create custom types for your keys
type authKey string
type logKey string

const safeAuthKey authKey = "id"
const safeLogKey logKey = "id"

func main() {
	bgCtx := context.Background()

	// ==========================================
	// SCENARIO A: THE DANGEROUS WAY (Using plain strings)
	// ==========================================
	fmt.Println("--- SCENARIO A: STRING COLLISION ---")

	// 1. Your Auth Middleware adds the User ID
	ctxDangerous := context.WithValue(bgCtx, "id", "User-Aakash")

	// 2. Later, a 3rd Party Logger adds its own ID using the same string
	ctxDangerous = context.WithValue(ctxDangerous, "id", "Transaction-999")

	// 3. Your database function tries to get the User ID back out...
	// DISASTER! It prints "Transaction-999" instead of "User-Aakash"
	fmt.Printf("Extracted User ID: %v\n\n", ctxDangerous.Value("id"))

	// ==========================================
	// SCENARIO B: THE SAFE WAY (Using custom types)
	// ==========================================
	fmt.Println("--- SCENARIO B: CUSTOM TYPES ---")

	// 1. Your Auth Middleware adds the User ID using your custom type
	ctxSafe := context.WithValue(bgCtx, safeAuthKey, "User-Aakash")

	// 2. The 3rd Party Logger adds its ID using ITS custom type
	// Even though both constants equal "id" underneath, Go treats them as completely different keys!
	ctxSafe = context.WithValue(ctxSafe, safeLogKey, "Transaction-999")

	// 3. Your database function safely extracts the correct User ID!
	fmt.Printf("Extracted User ID: %v\n", ctxSafe.Value(safeAuthKey))
	fmt.Printf("Extracted Logger ID: %v\n", ctxSafe.Value(safeLogKey))
}
