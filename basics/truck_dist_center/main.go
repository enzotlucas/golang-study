package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %+v\n", truck)

	time.Sleep(time.Second)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo, code: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo, code: %w", err)
	}

	fmt.Printf("Finished truck processing: %+v\n", truck)

	return nil
}

func processFleet(fleet []Truck) error {
	var wg sync.WaitGroup

	for _, t := range fleet {
		wg.Add(1)

		go func(Truck) {
			if err := processTruck(t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)
	}

	wg.Wait()

	return nil
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type EletricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *EletricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1

	return nil
}

func (e *EletricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1

	return nil
}

var (
	ErrNotImplemented = errors.New("NotImplemented")
	ErrTruckNotFound  = errors.New("TruckNotFound")
)

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&EletricTruck{id: "EL1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&EletricTruck{id: "EL2", cargo: 0, battery: 100},
	}

	if err := processFleet(fleet); err != nil {
		log.Fatalf("Error processing fleet: %+v", err)
	}
	log.Println("All trucks processed successfully!")
}
