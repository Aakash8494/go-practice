//go:build ignore

package main

import (
	"fmt"
	"time"
)

// A super simple waiter with just one basic loop
func simpleWaiter(pipe chan string) {
	// The waiter just stands here watching the pipe
	for order := range pipe {
		fmt.Println("👨‍🍳 Waiter grabbed:", order)
	}

	// This only prints AFTER the manager closes the pipe!
	fmt.Println("👨‍🍳 Waiter says: Pipe is closed, I am going home!")
}

func main() {
	// 1. Create a simple pipe
	ordersPipe := make(chan string, 5)

	// 2. Hire ONE waiter in the background
	go simpleWaiter(ordersPipe)

	// --- CASE 1: The First Order ---
	fmt.Println("👔 Manager: Dropping 'Pizza' into pipe...")
	ordersPipe <- "Pizza"
	time.Sleep(1 * time.Second) // Pausing so you can see the waiter react

	// --- CASE 2: The Late Order ---
	fmt.Println("\n👔 Manager: Dropping 'Burger' into pipe...")
	ordersPipe <- "Burger"
	time.Sleep(1 * time.Second)

	// --- CASE 3: Closing Time ---
	fmt.Println("\n👔 Manager: Closing the shop!")
	close(ordersPipe)

	time.Sleep(1 * time.Second) // Pausing to let the waiter say goodbye
}
