package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrTruckNotFound = errors.New("Truck not found")
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return ErrTruckNotFound
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo :%w", err)
	}
	return errors.New("Xato chiqdiku\n")
}

func main() {
	trucks := []Truck{
		{id: "First"},
		{id: "Second"},
		{id: "Third"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		err := processTruck(truck)
		if err != nil {
			log.Fatalf("Errore %s", err)
		}
	}
}
