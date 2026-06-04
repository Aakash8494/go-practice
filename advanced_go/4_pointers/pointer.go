//go:build ignore

package main

import "fmt"

type Truck struct {
	cargo int
}

// 1. PASS-BY-VALUE (The Bad Way)
// This receives a normal Truck. Go makes a complete COPY of it.
func updateCopy(t Truck) {
	t.cargo = 100
	// The instructor got an "unused write" warning here because
	// this copy is destroyed as soon as the function ends.
}

// 2. PASS-BY-REFERENCE (The Good Way)
// This receives a POINTER to a Truck (*Truck).
func updatePointer(t *Truck) {
	// 3. THE NIL CHECK
	// If someone passes a variable that hasn't been given an address yet,
	// trying to update it will cause a program crash (panic).
	if t == nil {
		fmt.Println("Error: Received a nil pointer!")
		return
	}

	t.cargo = 100 // This updates the ORIGINAL truck in place!
}

func main() {
	fmt.Println("--- PART 1: Memory Addresses & Basic Pointers ---")

	// A standard variable is just a box in memory.
	truckID := 42
	fmt.Printf("Value of truckID: %d\n", truckID)

	// The '&' operator gets the exact memory address of that box.
	fmt.Printf("Memory Address of truckID (&truckID): %p\n", &truckID)

	// Creating a pointer variable that holds the address of truckID.
	var anotherTruckID *int = &truckID
	fmt.Printf("Memory Address stored in anotherTruckID: %p\n", anotherTruckID)

	// The '*' operator "dereferences" the pointer. It travels to the address
	// and modifies the actual value inside the original box.
	*anotherTruckID = 0
	fmt.Printf("New Value of original truckID: %d\n\n", truckID) // This is now 0!

	fmt.Println("--- PART 2: Pointers with Structs ---")

	myTruck := Truck{cargo: 0}

	// Try to update using pass-by-value (Sends a clone)
	updateCopy(myTruck)
	fmt.Printf("Cargo after updateCopy: %d (Nothing changed!)\n", myTruck.cargo)

	// Try to update using a pointer (Sends the memory address using '&')
	updatePointer(&myTruck)
	fmt.Printf("Cargo after updatePointer: %d (Updated successfully!)\n\n", myTruck.cargo)

	fmt.Println("--- PART 3: The Danger of Nil Pointers ---")

	// If you declare a normal integer, it defaults to 0.
	var safeID int
	fmt.Printf("Default int: %d\n", safeID)

	// If you declare a pointer but don't give it an address, it defaults to 'nil'.
	var dangerousPointer *int
	fmt.Printf("Default pointer: %v\n", dangerousPointer)

	// If you uncommented the line below, the program would crash
	// ("invalid memory address") because it's trying to change data at an address that doesn't exist!
	// *dangerousPointer = 5
}
