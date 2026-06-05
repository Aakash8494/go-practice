//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. The Shared Map (Not concurrent safe!)
	scores := make(map[string]int)

	// 2. The WaitGroup (To stop the program from exiting early)
	var wg sync.WaitGroup

	// 3. Hire 100 workers and tell them to run at the same time
	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			// 🚨 THE CRASH HAPPENS HERE!
			// No mu.Lock() is used!
			// Everyone is trying to change "TeamA" simultaneously.
			scores["TeamA"] = workerID

		}(i)
	}

	// 4. Wait for everyone (but the program will crash before reaching here)
	wg.Wait()
	fmt.Println("Done! Final Score:", scores["TeamA"])
}
