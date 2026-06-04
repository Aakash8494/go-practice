//go:build ignore

package main

import (
	"fmt"
	"sync"
)

// A dummy function that returns an error
func processTruck(truckID int) error {
	return fmt.Errorf("engine failure on truck %d", truckID)
}

func main() {
	fleetOfTrucks := []int{1, 2, 3, 4, 5}

	// 'wg' renamed to 'taskTracker' so it's obvious what its job is
	var taskTracker sync.WaitGroup

	// 1. Create a BUFFERED Channel.
	// We tell it to hold exactly enough space for all our trucks (5).
	// This prevents Deadlocks!
	errorCollectionChannel := make(chan error, len(fleetOfTrucks))

	for _, currentTruck := range fleetOfTrucks {
		taskTracker.Add(1)

		// We pass 'currentTruck' into the Goroutine and call it 'activeTruckID'
		go func(activeTruckID int) {
			defer taskTracker.Done()

			processingError := processTruck(activeTruckID)

			if processingError != nil {
				// 2. Push the error INTO the channel pipe
				errorCollectionChannel <- processingError
			}
		}(currentTruck)
	}

	// 3. Wait for all background tasks to finish
	taskTracker.Wait()

	// 4. Close the pipe! You MUST do this when you are done sending,
	// otherwise the listener loop below will wait forever and cause a deadlock.
	close(errorCollectionChannel)

	// 5. Read all the errors OUT of the channel pipe
	// You can loop over a channel just like an array!
	fmt.Println("--- Error Report ---")
	for caughtError := range errorCollectionChannel {
		fmt.Println("Caught error:", caughtError)
	}
}
