package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

// The Interface (No locks mentioned here!)
type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error) // Returns a safe clone (Truck)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
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

	// Dereference pointer to return a safe clone
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

// WRITE Lock (Updating data)
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
