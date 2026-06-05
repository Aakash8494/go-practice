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
	m := make(map[string]int) // Keys MUST be strings, Values MUST be ints[cite: 16, 17, 20].

	// Adding data to the map
	m["A"] = 1
	m["B"] = 2

	// 2. Reading and Checking if a Key Exists
	// Go returns TWO things when you ask for a value: the value itself, and a boolean (usually called 'ok' or 'exists')[cite: 28, 29].
	val, ok := m["A"]
	if ok {
		fmt.Println("Key A exists! Value:", val) // Safe to proceed! [cite: 30]
	}

	// 3. The "Inline" Check (Very common in Go)
	// You check the map and write the 'if' statement all on one line[cite: 31, 32].
	if valB, exists := m["B"]; exists {
		fmt.Println("Key B exists! Value:", valB)
	}

	// 4. Deleting a Key
	// Use the 'delete' keyword. Pass the map first, then the key you want to delete[cite: 36, 37].
	delete(m, "A")

	// 5. Clearing the whole Map
	// Wipes everything in the map completely[cite: 39].
	clear(m)

	// 6. The "maps" package tools
	// Go gives you built-in tools to clone a map or securely check if two maps are exactly equal[cite: 42, 43].
	map1 := map[string]int{"X": 10}
	map2 := maps.Clone(map1)
	isEqual := maps.Equal(map1, map2)
	fmt.Println("Are the maps equal?", isEqual)
}
