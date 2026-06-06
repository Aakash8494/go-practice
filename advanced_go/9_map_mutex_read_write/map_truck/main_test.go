package main_truck

import (
	"testing"
)

// --- BASIC TESTS ---

func TestAddTruck(t *testing.T) {
	m := NewTruckManager()
	err := m.AddTruck("1", 100)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetTruck(t *testing.T) {
	m := NewTruckManager()
	m.AddTruck("1", 100)

	truck, err := m.GetTruck("1")
	if err != nil || truck.Cargo != 100 {
		t.Errorf("Failed to safely get truck")
	}
}

func TestRemoveTruck(t *testing.T) {
	m := NewTruckManager()
	m.AddTruck("1", 100)

	m.RemoveTruck("1")
	_, err := m.GetTruck("1")

	if err != ErrTruckNotFound {
		t.Errorf("Expected ErrTruckNotFound, got %v", err)
	}
}

func TestUpdateTruckCargo(t *testing.T) {
	m := NewTruckManager()
	m.AddTruck("1", 100)

	m.UpdateTruckCargo("1", 500)
	truck, _ := m.GetTruck("1")

	if truck.Cargo != 500 {
		t.Errorf("Expected cargo 500, got %d", truck.Cargo)
	}
}

// --- THE BIG CONCURRENCY TEST ---

func TestConcurrentUpdate(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)

	const numGoroutines = 100
	const iterations = 100
	done := make(chan bool)

	// Launch 100 workers
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				// 🛡️ THE FIX: We use our new atomic function so the math is safe!
				manager.AddCargo("1", 1)
			}
			done <- true
		}()
	}

	// Wait for all 100 workers
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	expectedFinalValue := numGoroutines*iterations + 100
	finalTruck, _ := manager.GetTruck("1")

	if finalTruck.Cargo != expectedFinalValue {
		t.Errorf("Expected %d but got %d", expectedFinalValue, finalTruck.Cargo)
	}
}
