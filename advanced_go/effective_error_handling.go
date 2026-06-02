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

// ---------------------------------------------------------
// THE INTERFACE
// ---------------------------------------------------------
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

// ---------------------------------------------------------
// TYPE 1: NormalTruck (What used to just be "Truck")
// ---------------------------------------------------------
type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	if t.id == "Truck-3" {
		return ErrTruckNotFound
	}
	if t.id == "Truck-2" {
		return ErrCargoDoorJammed
	}
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	if t.id == "Truck-1" {
		return ErrCargoDoorJammed
	}
	t.cargo = 0
	return nil
}

// ---------------------------------------------------------
// TYPE 2: ElectricTruck
// ---------------------------------------------------------
type ElectricTruck struct {
	id      string
	cargo   int
	battery int
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

// ---------------------------------------------------------
// THE PROCESSOR
// ---------------------------------------------------------
// processTruck receives the INTERFACE. It doesn't care which struct it gets!
func processTruck(truck Truck) error {
	// We can't print truck.id directly here anymore because the interface
	// doesn't know about "id", it only knows about the two methods.
	// But it works perfectly for calling the methods!

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	// Because our methods use pointer receivers like `(t *NormalTruck)`,
	// we use the '&' symbol to pass their memory addresses into the slice.
	trucks := []Truck{
		&NormalTruck{id: "Truck-1"},
		&NormalTruck{id: "Truck-2"},
		&NormalTruck{id: "Truck-3"},
		&ElectricTruck{id: "EV-1", battery: 100}, // Mixed in the same array!
	}

	// The new detailed loop with error matching
	for _, truck := range trucks {
		fmt.Printf("Processing: %+v\n", truck)

		err := processTruck(truck)

		if err != nil {
			if errors.Is(err, ErrTruckNotFound) {
				fmt.Printf("ALERT: We lost it! It is not in the system.\n")
			} else if errors.Is(err, ErrCargoDoorJammed) {
				fmt.Printf("MAINTENANCE: Send mechanic. %v\n", err)
			} else {
				fmt.Printf("Error: Could not process because %v\n", err)
			}
		} else {
			fmt.Println("Processed successfully.")
		}
		fmt.Println("---")
	}

	fmt.Println("-----------------------")
	fmt.Println("-----------------------")

	// Your original testing loop, completely untouched!
	// (It will just run through the same trucks array a second time)
	for _, truck := range trucks {
		if err := processTruck(truck); err != nil {
			// log.Fatalf("Error processing truck: %s", err)
			fmt.Println("Error processing truck: Println --->", err)
			fmt.Printf("Error processing truck: Printf ---> %s\n", err)
		}
	}
}
