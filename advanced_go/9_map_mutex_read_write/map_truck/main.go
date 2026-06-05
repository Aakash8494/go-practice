package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

// The Interface (Updated to include our new AddCargo tool)
type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
	AddCargo(id string, amount int) error // <-- NEW RULE
}

type Truck struct {
	ID    string
	Cargo int
}

// The Manager with embedded Read-Write Lock
type truckManager struct {
	sync.RWMutex
	trucks map[string]*Truck
}

// Constructor to wake up the map
func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

// WRITE Lock (Adding data)
func (m *truckManager) AddTruck(id string, cargo int) error {
	m.Lock()
	defer m.Unlock()

	m.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

// READ Lock (Fetching data)
func (m *truckManager) GetTruck(id string) (Truck, error) {
	m.RLock()
	defer m.RUnlock()

	truck, ok := m.trucks[id]
	if !ok {
		return Truck{}, ErrTruckNotFound
	}

	return *truck, nil
}

// WRITE Lock (Deleting data)
func (m *truckManager) RemoveTruck(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.trucks[id]; !ok {
		return ErrTruckNotFound
	}

	delete(m.trucks, id)
	return nil
}

// WRITE Lock (Updating exact data)
func (m *truckManager) UpdateTruckCargo(id string, cargo int) error {
	m.Lock()
	defer m.Unlock()

	truck, ok := m.trucks[id]
	if !ok {
		return ErrTruckNotFound
	}

	truck.Cargo = cargo
	return nil
}

// 🛡️ THE FIX: Atomic Read-Modify-Write
func (m *truckManager) AddCargo(id string, amount int) error {
	m.Lock()         // 1. LOCK THE DOOR
	defer m.Unlock() // 4. UNLOCK WHEN FINISHED

	truck, ok := m.trucks[id]
	if !ok {
		return ErrTruckNotFound
	}

	// 2 & 3. READ AND WRITE THE MATH WHILE THE DOOR IS STILL LOCKED!
	truck.Cargo = truck.Cargo + amount
	return nil
}
