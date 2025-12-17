package main

//our goal here is to
//implement an interface for managing a fleet of trucks
//with methods to add, get, remove, and update trucks.
//we will also define a struct to represent a truck
//and a struct to manage the collection of trucks.

import (
	"errors"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *truckManager) AddTruck(id string, cargo int) error {
	tm.trucks[id] = &Truck{
		ID:    id,
		Cargo: cargo,
	}
	return nil
}

func (tm *truckManager) GetTruck(id string) (*Truck, error) {
	truck, exists := tm.trucks[id]
	if !exists {
		return nil, ErrTruckNotFound
	}
	return truck, nil
}

func (tm *truckManager) RemoveTruck(id string) error {
	_, exists := tm.trucks[id]
	if !exists {
		return ErrTruckNotFound
	}
	delete(tm.trucks, id)
	return nil
}

func (tm *truckManager) UpdateTruckCargo(id string, cargo int) error {
	truck, exists := tm.trucks[id]
	if !exists {
		return ErrTruckNotFound
	}
	truck.Cargo = cargo
	return nil
}

func main() {
	// This is just a placeholder to make the package "main" complete.
}
