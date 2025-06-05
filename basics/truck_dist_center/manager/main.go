package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")
var ErrTruckAlreadyExists = errors.New("truck already exists")
var ErrInvalidRequest = errors.New("invalid request")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
	sync.RWMutex
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *truckManager) AddTruck(id string, cargo int) error {
	tm.Lock()
	defer tm.Unlock()

	if id == "" || cargo < 0 {
		return ErrInvalidRequest
	}
	if _, exists := tm.trucks[id]; exists {
		return ErrTruckAlreadyExists
	}
	tm.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (tm *truckManager) GetTruck(id string) (Truck, error) {
	tm.RLock()
	defer tm.RUnlock()

	if id == "" {
		return Truck{}, ErrInvalidRequest
	}
	t, exists := tm.trucks[id]
	if !exists {
		return Truck{}, ErrTruckNotFound
	}
	return *t, nil
}

func (tm *truckManager) RemoveTruck(id string) error {
	tm.Lock()
	defer tm.Unlock()

	if id == "" {
		return ErrInvalidRequest
	}
	if _, exists := tm.trucks[id]; !exists {
		return ErrTruckNotFound
	}
	delete(tm.trucks, id)
	return nil
}

func (tm *truckManager) UpdateTruckCargo(id string, cargo int) error {
	tm.Lock()
	defer tm.Unlock()

	if id == "" || cargo < 0 {
		return ErrInvalidRequest
	}
	if _, exists := tm.trucks[id]; !exists {
		return ErrTruckNotFound
	}
	tm.trucks[id].Cargo = cargo
	return nil
}
