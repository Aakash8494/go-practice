//go:build ignore

package main

import (
	"fmt"
	"sort"
	"strings"
)

func TopAnomalies(rawLogs string) []string {
	// 1. Break the massive string into individual lines
	lines := strings.Split(rawLogs, "\n")

	// 2. Initialize a map to count the frequencies of the cleaned errors
	anomalyCounts := make(map[string]int)

	// 3. Process each line
	for _, line := range lines {
		// Quick safety check: ignore empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Check for the exclusion keyword FIRST to save processing time
		if strings.Contains(line, "RETRY") {
			continue
		}

		var message string

		// Filter by prefix and extract the core message
		if strings.HasPrefix(line, "CRITICAL:") {
			message = strings.TrimPrefix(line, "CRITICAL:")
		} else if strings.HasPrefix(line, "FATAL:") {
			message = strings.TrimPrefix(line, "FATAL:")
		} else {
			// If it's INFO, WARN, or anything else, ignore it
			continue
		}

		// 4. Clean the extracted message
		// First, strip the leading/trailing spaces
		message = strings.TrimSpace(message)
		// Next, strip any trailing punctuation specifically requested
		message = strings.TrimRight(message, ".!?")
		// Clean any remaining spaces that might have been hiding behind the punctuation
		message = strings.TrimSpace(message)

		// 5. Aggregate the frequency
		anomalyCounts[message]++
	}

	// 6. Sort the results
	// Go maps cannot be sorted directly, so we must transfer the data to a slice of structs
	type kv struct {
		Key   string
		Value int
	}

	var sortedAnomalies []kv
	for k, v := range anomalyCounts {
		sortedAnomalies = append(sortedAnomalies, kv{k, v})
	}

	// Sort the slice in descending order based on the Value (frequency)
	sort.Slice(sortedAnomalies, func(i, j int) bool {
		return sortedAnomalies[i].Value > sortedAnomalies[j].Value
	})

	// 7. Extract the top 3 (or fewer if there aren't 3 unique items)
	var top3 []string
	for i, kv := range sortedAnomalies {
		if i == 3 {
			break
		}
		top3 = append(top3, kv.Key)
	}

	return top3
}

func main() {
	rawLogs := `
INFO: System boot successful.
CRITICAL:   Database connection timeout   
WARN: Memory usage at 75%
FATAL: Segmentation fault in worker thread!
CRITICAL: Database connection timeout.
INFO: Request served in 20ms
CRITICAL: Cache miss on user profile. RETRY initiated.
FATAL:   Segmentation fault in worker thread   ?
CRITICAL: Disk space critically low!
CRITICAL: Database connection timeout!
`
	result := TopAnomalies(rawLogs)

	fmt.Println("--- Top Anomalies ---")
	for i, anomaly := range result {
		fmt.Printf("%d. %s\n", i+1, anomaly)
	}
}
