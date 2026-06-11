//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. Our shared data (The Bathroom)
	scores := make(map[string]int)

	// 2. Our Bouncer (To keep the program alive)
	var wg sync.WaitGroup

	// 3. Our Door Lock (To prevent crashes)
	var mu sync.Mutex

	// Let's hire 5 workers (Goroutines) to add score at the exact same time
	for i := 1; i <= 5; i++ {

		wg.Add(1) // Tell bouncer: "One more worker starting!"

		go func(workerID int) {
			defer wg.Done() // Tell bouncer: "Worker finished!" when done

			// 🚪 WAIT IN LINE AND LOCK THE DOOR
			mu.Lock()

			// --- WE ARE INSIDE THE ROOM ---
			// Only one worker can be on this line of code at a time!
			scores["TeamA"] += 10
			fmt.Printf("Worker %d added points! Total is now: %d\n", workerID, scores["TeamA"])
			// ------------------------------

			// 🔓 UNLOCK THE DOOR FOR THE NEXT GUY
			mu.Unlock()

		}(i)
	}

	// 4. Wait for all 5 workers to finish going through the door
	wg.Wait()
	fmt.Println("All workers done! Final Score:", scores["TeamA"])
}
