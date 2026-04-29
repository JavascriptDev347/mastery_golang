package main

import (
	"fmt"
	"time"
)

func sayHello(msg string) {
	fmt.Println("Hello " + msg)
}

func main() {

	start := time.Now()
	// call the function in a goroutine
	go sayHello("Name 1")
	go sayHello("Name 2")
	go sayHello("Name 3")

	sayHello("Russl")
	sayHello("Bussel")
	time.Sleep(2 * time.Millisecond)
	end := time.Now()

	fmt.Printf("Execution time: %s\n", end.Sub(start))
	fmt.Println("Main function completed")
}
