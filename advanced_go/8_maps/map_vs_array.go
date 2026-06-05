//go:build ignore

package main

import "fmt"

func main() {
	// THE ARRAY WAY (Requires looping)
	// We want to find "Rahul", but we don't know his index.
	namesArr := []string{"Amit", "Priya", "Rahul", "Neha"}

	fmt.Println("=== THE ARRAY WAY ===")
	for index, name := range namesArr {
		fmt.Printf("Checking index %d... is it Rahul? It's %s\n", index, name)

		if name == "Rahul" {
			fmt.Println("-> Found Rahul in Array! (But we had to check the others first)")
			break // We stop the loop since we found him!
		}
	}

	// THE MAP WAY (Instant lookup)
	// We ask for "Rahul" directly using his key.
	namesMap := map[string]bool{
		"Amit":  true,
		"Priya": true,
		"Rahul": true,
		"Neha":  true,
	}

	fmt.Println("\n=== THE MAP WAY ===")
	fmt.Println("Asking the map directly for 'Rahul' without any loops...")

	// No loop needed! Just ask the map.
	if namesMap["Rahul"] {
		fmt.Println("-> Found Rahul in Map! (Instantly in 1 step!)")
	}
}
