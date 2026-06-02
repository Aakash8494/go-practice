package main

import (
	"testing"
)

// The function name must start with "Test"
// It receives a pointer to the testing suite: (t *testing.T)
func TestProcessTruck(t *testing.T) {
	// 1. Setup our concrete test subjects
	normalTruck := &NormalTruck{id: "Test-Normal", cargo: 10} // Pre-load with 10 cargo
	electricTruck := &ElectricTruck{id: "Test-EV", cargo: 0, battery: 100}

	// 2. Run the business logic we want to test
	// We can pass both into processTruck because they satisfy the Truck interface!
	_ = processTruck(normalTruck)
	_ = processTruck(electricTruck)

	// 3. Assertions (The Go Way)

	// Test A: Did the normal truck get unloaded properly?
	// It should be 0, because processTruck loads 1, then unloads everything to 0.
	if normalTruck.cargo != 0 {
		// t.Fatalf automatically fails the test and prints this error
		t.Fatalf("Normal truck cargo should be 0, but got %d", normalTruck.cargo)
	}

	// Test B: Did the electric truck use exactly 2 battery?
	// Started at 100. Load uses 1. Unload uses 1. Should be 98.
	// (Note: In the video, his logic just blindly decreased by 1 twice, resulting in -2 because he started at 0).
	if electricTruck.battery != 98 {
		t.Fatalf("Electric truck battery should be 98, but got %d", electricTruck.battery)
	}
}
