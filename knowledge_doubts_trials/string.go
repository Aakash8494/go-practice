//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Go 'strings' Package: Daily Driver Functions ===")

	// ---------------------------------------------------------
	// 1. Inspection and Searching (Checking what a string holds)
	// ---------------------------------------------------------
	fmt.Println("\n--- 1. Inspecting & Searching ---")

	statement := "Go is fast and fun."

	// HasSuffix: Checks if a string ends with a specific substring (Your example!)
	fmt.Printf("HasSuffix: %v\n", strings.HasSuffix(statement, ".")) // Output: true

	// HasPrefix: Checks if a string starts with a specific substring
	fmt.Printf("HasPrefix: %v\n", strings.HasPrefix(statement, "Go")) // Output: true

	// Contains: Checks if a substring exists anywhere in the string
	fmt.Printf("Contains:  %v\n", strings.Contains(statement, "fast")) // Output: true

	// Index: Returns the starting position of a substring, or -1 if not found
	fmt.Printf("Index:     %v\n", strings.Index(statement, "is")) // Output: 3

	// Count: Counts non-overlapping instances of a substring
	fmt.Printf("Count:     %v\n", strings.Count(statement, "f")) // Output: 2

	// ---------------------------------------------------------
	// 2. Trimming and Cleaning (Fixing messy inputs)
	// ---------------------------------------------------------
	fmt.Println("\n--- 2. Trimming & Cleaning ---")

	messyInput := "   \t  user_input_with_whitespace \n  "

	// TrimSpace: Removes all leading and trailing whitespace (crucial for API/Form inputs)
	fmt.Printf("TrimSpace:  '%v'\n", strings.TrimSpace(messyInput))

	// TrimSuffix/TrimPrefix: Great for removing known file extensions or tags
	filename := "report.csv"
	fmt.Printf("TrimSuffix: '%v'\n", strings.TrimSuffix(filename, ".csv")) // Output: 'report'

	// ---------------------------------------------------------
	// 3. Splitting and Joining (Converting between Strings and Slices)
	// ---------------------------------------------------------
	fmt.Println("\n--- 3. Splitting & Joining ---")

	csvLine := "apple,banana,cherry"

	// Split: Breaks a string into a slice of strings based on a separator
	fruits := strings.Split(csvLine, ",")
	fmt.Printf("Split:      %q\n", fruits) // %q prints the slice with quotes: ["apple" "banana" "cherry"]

	// Join: Combines a slice of strings into a single string with a separator
	joinedFruits := strings.Join(fruits, " | ")
	fmt.Printf("Join:       %v\n", joinedFruits) // Output: apple | banana | cherry

	// ---------------------------------------------------------
	// 4. Replacing and Modifying (Changing content)
	// ---------------------------------------------------------
	fmt.Println("\n--- 4. Replacing & Modifying ---")

	typo := "foo bar foo"

	// ReplaceAll: Replaces every instance of the old string with the new one
	fmt.Printf("ReplaceAll: %v\n", strings.ReplaceAll(typo, "foo", "baz")) // Output: baz bar baz

	// Replace: Replaces a specific number of instances (n). Use -1 for all.
	fmt.Printf("Replace(1): %v\n", strings.Replace(typo, "foo", "baz", 1)) // Output: baz bar foo

	// ---------------------------------------------------------
	// 5. Casing (Formatting text)
	// ---------------------------------------------------------
	fmt.Println("\n--- 5. Casing ---")

	fmt.Printf("ToUpper:    %v\n", strings.ToUpper("shouting"))     // Output: SHOUTING
	fmt.Printf("ToLower:    %v\n", strings.ToLower("QUIET PLEASE")) // Output: quiet please
}
