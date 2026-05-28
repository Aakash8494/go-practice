//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func reverseWordsWithPeriod(statement string) string {
	hasPeriod := false

	// Step 1: Check for and remove the period at the end
	if strings.HasSuffix(statement, ".") {
		hasPeriod = true
		statement = strings.TrimSuffix(statement, ".")
	}

	// Step 2: Split the sentence into words
	words := strings.Fields(statement)
	// words := strings.Split(statement, " ")

	// Step 3: Swap the words from start to end
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Step 4: Join them back with a space
	result := strings.Join(words, " ")

	// Step 5: Add the period back to the end if it was there originally
	if hasPeriod {
		result += "."
	}

	return result
}

func main() {
	text := "hello from the other side."
	fmt.Println(reverseWordsWithPeriod(text))
	// Output: side other the from hello.
}
