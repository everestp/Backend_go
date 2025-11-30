package main

import (
	"errors"
	"fmt"
)

// ------------------------------
// Error definitions
// ------------------------------

var (
	// ErrNotImplemented indicates functionality that has not yet been written.
	ErrNotImplemented = errors.New("not implemented")

	// ErrTruckNotFound represents a failure to locate a truck or cargo.
	ErrTruckNotFound = errors.New("truck not found")
)

// ------------------------------
// Truck Interface
// ------------------------------

// Truck represents a generic vehicle capable of loading and unloading cargo.
// Go interfaces define *behavior*, not data.
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
	ID() string
}

// ------------------------------
// NormalTruck implementation
// ------------------------------

// NormalTruck is a concrete implementation of Truck.
type NormalTruck struct {
	id string
}

// ID returns the unique identifier of the truck.
func (t *NormalTruck) ID() string {
	return t.id
}

// LoadCargo simulates loading cargo onto the truck.
func (t *NormalTruck) LoadCargo() error {
	// Simulate error to show error handling concepts.
	return fmt.Errorf("failed to load cargo: %w", ErrTruckNotFound)
}

// UnloadCargo simulates unloading cargo.
func (t *NormalTruck) UnloadCargo() error {
	// No error means unload successful.
	return nil
}

// ------------------------------
// Truck Processing Logic
// ------------------------------

// processTruck handles loading + unloading for a single truck.
// Demonstrates:
// - Interface usage
// - Error wrapping with %w
// - Clean structured logging style
func processTruck(t Truck) error {
	fmt.Printf("[INFO] Processing truck: %s\n", t.ID())

	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("[ERROR] unable to load cargo for %s: %w", t.ID(), err)
	}

	if err := t.UnloadCargo(); err != nil {
		return fmt.Errorf("[ERROR] unable to unload cargo for %s: %w", t.ID(), err)
	}

	fmt.Printf("[OK] Completed processing: %s\n", t.ID())
	return nil
}

// ------------------------------
// Fleet Processing
// ------------------------------

// processFleet processes multiple trucks.
// Shows:
// - slice iteration
// - error continuation control
// - separation of concerns
// - clean logging style
func processFleet(fleet []Truck) {
	fmt.Println("=== Fleet Processing Start ===")

	for _, truck := range fleet {
		fmt.Printf("[ARRIVED] %s\n", truck.ID())

		err := processTruck(truck)
		if err != nil {
			fmt.Printf("[WARN] %s skipped due to error: %v\n", truck.ID(), err)
			continue // Try next truck
		}
	}

	fmt.Println("=== Fleet Processing Complete ===")
}

// ------------------------------
// main
// ------------------------------

func main() {
	// NOTE: Use pointers because the interface has pointer receivers.
	fleet := []Truck{
		&NormalTruck{id: "Truck-1"},
		&NormalTruck{id: "Truck-2"},
		&NormalTruck{id: "Truck-3"},
	}

	processFleet(fleet)
}
