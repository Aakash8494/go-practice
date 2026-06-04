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
// TYPE 1: NormalTruck
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
func processTruck(truck Truck) error {
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	trucks := []Truck{
		&NormalTruck{id: "Truck-1"},
		&NormalTruck{id: "Truck-2"},
		&NormalTruck{id: "Truck-3"},
		&ElectricTruck{id: "EV-1", battery: 100},
	}

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
	fmt.Println("--- NEW: EMPTY INTERFACES & TYPE ASSERTION ---")
	fmt.Println("-----------------------")

	// From the screenshot: Creating a map that accepts ANY value type
	person := make(map[string]interface{}, 0)
	person["name"] = "Tiago" // Storing a string
	person["age"] = 42       // Storing an int

	// Type Assertion:
	// 1. Try to grab "width"
	// 2. Assert it is an integer by adding `.(int)`
	// 3. Go returns the value AND a boolean (`exists`) letting us know if it worked
	val, exists := person["width"].(int)

	if !exists {
		// I changed log.Fatal to fmt.Println so it doesn't kill your app!
		fmt.Println("Error: width does not exist or is not an integer")
	} else {
		fmt.Println("Value found:", val)
	}
}

// =====================================================================
// --- LECTURE LEARNINGS (Cheat Sheet) ---
// =====================================================================
/*
1. Abstraction over Concreteness:
   Rely on abstractions (interfaces) rather than concrete implementations (structs).

2. Interfaces as Blueprints:
   An interface is essentially just a blueprint of method signatures.

3. Automatic/Implicit Implementation:
   Go does not use an 'implements' keyword. If a struct satisfies the interface by having the exact methods required, the compiler automatically accepts it.

4. Pointer Receivers vs. Value Receivers:
   To update a struct's field (like `cargo += 1`) inside a method, you MUST use a pointer receiver (`*NormalTruck`). Otherwise, you are just modifying a temporary clone.

5. Printing Struct Fields:
   Use `%+v` in printf to see the property names (like `battery:100`) alongside their values.

6. The `error` interface:
   The built-in `error` type is an interface. You create custom errors by implementing it.

7. Error String Formatting:
   Error strings should not be capitalized.

8. Empty Interfaces (`interface{}` / `any`):
   An empty interface specifies zero methods, meaning any type (string, struct, error) can satisfy it. In Go 1.18+, use `any`, which is just an alias for `interface{}`.

9. Type Assertion (From the screenshot!):
   When you pull a value out of an empty interface, the compiler doesn't know what type it is. You must evaluate it using "Type Assertion" by appending `.(type)` to it (e.g., `person["width"].(int)`). This returns the value and a boolean confirming if the assertion was successful.
*/
