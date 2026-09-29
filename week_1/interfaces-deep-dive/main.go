package main

import (
	"fmt"
)

type Truck interface {
	LoadCargo() error
	UnLoadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t NormalTruck) LoadCargo() error {
	return nil
}

func (t NormalTruck) UnLoadCargo() error {
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t ElectricTruck) LoadCargo() error {
	return nil
}

func (t ElectricTruck) UnLoadCargo() error {
	return nil
}

func processTruck(truck Truck) error {
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("failed to load cargo: %w", err)
	}
	return nil
}

func main2() {

	// trucks := []NormalTruck{
	// 	{id: "First", cargo: 12},
	// 	{id: "Second", cargo: 8},
	// }

	// eTrucks := []ElectricTruck{
	// 	{id: "First e", cargo: 10, battery: 50},
	// 	{id: "Second e", cargo: 5, battery: 30},
	// }

	// err := processTruck(NormalTruck{id: "1"})
	// if err != nil {
	// 	log.Fatalf("Error processing truck %s: %v", err)
	// }

	// err = processTruck(ElectricTruck{id: "2"})
	// if err != nil {
	// 	log.Fatalf("Error processing truck %s: %v", err)
	// }
	//
}
