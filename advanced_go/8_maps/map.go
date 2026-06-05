//go:build ignore

package main

import (
	"fmt"
	"maps"
)

func main() {
	// 1. Creating a Map
	// You use the 'make' keyword.
	// syntax: map[KeyType]ValueType
	m := make(map[string]int) // Keys MUST be strings, Values MUST be ints.
	fmt.Println("Step 0 -> After declaring the map variable :", m)

	// Adding data to the map
	m["A"] = 1
	m["B"] = 2
	fmt.Println("Step 1 -> After adding data:", m)

	// 2. Reading and Checking if a Key Exists
	// Go returns TWO things when you ask for a value: the value itself, and a boolean
	val, ok := m["A"]
	if ok {
		fmt.Println("Step 2 -> Key A exists! Value:", val)
	}

	// 3. The "Inline" Check (Very common in Go)
	// You check the map and write the 'if' statement all on one line.
	if valB, exists := m["B"]; exists {
		fmt.Println("Step 3 -> Key B exists! Value:", valB)
	}

	// 4. Deleting a Key
	// Use the 'delete' keyword. Pass the map first, then the key you want to delete.
	delete(m, "A")
	fmt.Println("Step 4 -> After deleting key 'A':", m)

	// 5. Clearing the whole Map
	// Wipes everything in the map completely.
	clear(m)
	fmt.Println("Step 5 -> After clearing the map completely:", m)

	// 6. The "maps" package tools
	// Go gives you built-in tools to clone a map or securely check if two maps are exactly equal.
	map1 := map[string]int{"X": 10}
	map2 := maps.Clone(map1)
	fmt.Println("Step 6 -> Map 1 is:", map1, "and Map 2 is:", map2)

	isEqual := maps.Equal(map1, map2)
	fmt.Println("Are the maps equal?", isEqual)
}
