package main

import (
	"errors"
	"fmt"
)

type Truck struct {
	id string
}

// processTruck handles the loading and unloading of a truck.
// Now it returns an 'error' type.
func processTruck(truck Truck) error {
	// Let's simulate an error for Truck-2 just to test it
	if truck.id == "Truck-2" {
		return errors.New("cargo door jammed")
	}

	fmt.Printf("Processing truck: %s\n", truck.id)
	return nil // nil means no error happened
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)

		// Catch the error returned by the function
		err := processTruck(truck)

		// Check if there is an actual error
		if err != nil {
			fmt.Printf("Error: Could not process %s because %v\n", truck.id, err)
		} else {
			fmt.Println("Processed successfully.")
		}
		fmt.Println("---")
	}
}
