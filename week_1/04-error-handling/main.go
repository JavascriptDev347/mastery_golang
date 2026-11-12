package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s, %s", e.Field, e.Message)
}

func createUser(name string, age int) error {
	if name == "" {
		return &ValidationError{
			Field:   "name",
			Message: "Name should not be empty !!!",
		}
	}

	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "age should be greater or equal 1",
		}
	}

	return nil
}

func register(name string, age int) error {
	if err := createUser(name, age); err != nil {
		return fmt.Errorf("foydalanuvchini ro‘yxatdan o‘tkazish: %w", err)
	}

	return nil
}

func main() {
	name, age := "", 21
	err := register(name, age)
	if err == nil {
		return
	}
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("%s maydonida xato: %s\n", validationErr.Field, validationErr.Message)
		return
	}
	fmt.Println("Kutilmagan xato:", err)
}

// 	trucks := []Truck{
// 		{id: "First"},
// 		{id: "Second"},
// 		{id: "Third"},
// 	}
//
// 	for _, truck := range trucks {
// 		fmt.Printf("Truck %s arrived.\n", truck.id)
// 		err := processTruck(truck)
// 		if err != nil {
// 			log.Fatalf("Errore %s", err)
// 		}
// 	}
// var (
// 	ErrTruckNotFound = errors.New("Truck not found")
// )
//
// type Truck struct {
// 	id string
// }
//
// func (t *Truck) LoadCargo() error {
// 	return ErrTruckNotFound
// }
//
// func processTruck(truck Truck) error {
// 	fmt.Printf("Processing truck: %s\n", truck.id)
// 	if err := truck.LoadCargo(); err != nil {
// 		return fmt.Errorf("Error loading cargo :%w", err)
// 	}
// 	return errors.New("Xato chiqdiku\n")
// }
