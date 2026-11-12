package main

import "fmt"

func main() {
	var p *int = nil

	var i interface{} = p

	fmt.Printf("Value: %v\n", i)
	fmt.Printf("Type: %T\n", i)
	fmt.Printf("Is nil: %v\n", i == nil)
}
