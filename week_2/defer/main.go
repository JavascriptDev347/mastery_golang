package main

import "fmt"

func steps() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("Asosiy ish bajarildi")
}

func main() {
	steps()
}
