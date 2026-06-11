//go:build ignore

package main

import (
	"fmt"
	"sync"
)

// 1. A simple Truck struct
type Truck struct {
	ID    string
	Cargo int
}

// 2. The Manager with the embedded RWMutex
type truckManager struct {
	// We embed the lock directly so we can use m.Lock()[cite: 109, 110].
	sync.RWMutex
	trucks map[string]*Truck
}

// 3. Adding a Truck (Uses a WRITE Lock)
func (m *truckManager) AddTruck(id string, cargo int) {
	m.Lock()         // 🚪 Lock door for EVERYONE (Writers & Readers)
	defer m.Unlock() // 🔓 Unlock when done

	m.trucks[id] = &Truck{ID: id, Cargo: cargo}
}

// 4. Getting a Truck (Uses a READ Lock & Dereferencing)
// Notice it returns 'Truck' (a clone), NOT '*Truck'[cite: 170, 171].
func (m *truckManager) GetTruck(id string) (Truck, error) {
	// 📖 Read Lock: Multiple readers can enter at once[cite: 142, 143]!
	m.RLock()

	// 📖 Unlock reader door when the function finishes[cite: 139].
	defer m.RUnlock()

	truck, exists := m.trucks[id]
	if !exists {
		// Return an empty truck and an error if missing
		return Truck{}, fmt.Errorf("truck not found")
	}

	// 🚨 THE FIX: We use '*' to dereference the pointer[cite: 172, 173]!
	// This makes a safe clone of the truck so Goroutines don't crash.
	return *truck, nil
}

func main() {
	// Initialize our manager and the empty map
	m := &truckManager{
		trucks: make(map[string]*Truck),
	}

	// Safely Add a truck
	m.AddTruck("Truck-1", 100)

	// Safely Get a truck clone
	t, err := m.GetTruck("Truck-1")
	if err == nil {
		fmt.Printf("Got truck safely! ID: %s, Cargo: %d\n", t.ID, t.Cargo)
	}
}
