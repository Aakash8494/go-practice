package main

import (
	"errors"
	"fmt"
)

// 1. Sentinel Errors
var (
	ErrNotImplemented  = errors.New("not implemented")
	ErrTruckNotFound   = errors.New("truck not found")
	ErrCargoDoorJammed = errors.New("cargo door jammed")
)

type Truck struct {
	id string
}

// A method attached to the truck
func (t *Truck) LoadCargo() error {
	if t.id == "Truck-3" {
		return ErrTruckNotFound
	}
	if t.id == "Truck-2" {
		return ErrCargoDoorJammed
	}
	return nil
}

// NEW: The UnloadCargo method (From the screenshot)
func (t *Truck) UnloadCargo() error {
	if t.id == "Truck-1" {
		return ErrCargoDoorJammed
	}
	return nil
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	// Loading Phase
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	// NEW: Unloading Phase (From the screenshot)
	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error unloading cargo: %w", err)
	}

	return nil // Changed from ErrNotImplemented to nil to allow success!
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
		{id: "Truck-4"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)

		// Catch the error returned by the function
		err := processTruck(truck)

		if err != nil {
			// ERROR MATCHING: Still works perfectly with both Load and Unload errors!
			if errors.Is(err, ErrTruckNotFound) {
				fmt.Printf("ALERT: We lost %s! It is not in the system.\n", truck.id)
			} else if errors.Is(err, ErrCargoDoorJammed) {
				fmt.Printf("MAINTENANCE: Send mechanic to %s. %v\n", truck.id, err)
			} else {
				// A fallback for any unknown errors
				fmt.Printf("Error: Could not process %s because %v\n", truck.id, err)
			}
		} else {
			fmt.Println("Processed successfully.")
		}
		fmt.Println("---")
	}

	fmt.Println("-----------------------")
	fmt.Println("-----------------------")

	// Your original testing loop, completely untouched!
	// for _, truck := range trucks {
	// 	if err := processTruck(truck); err != nil {
	// 		// log.Fatalf("Error processing truck: %s", err)
	// 		fmt.Println("Error processing truck: Println --->", err)
	// 		fmt.Printf("Error processing truck: Printf ---> %s\n", err)
	// 	}
	// }
}
