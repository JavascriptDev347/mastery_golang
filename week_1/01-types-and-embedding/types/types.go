package typesandembedding

import "fmt"

// what is an embedding?
// embedding is a way to include one struct within another struct.
// It allows us to reuse code and create more complex data structures.
//  The embedded struct can be accessed directly from the outer struct, and its fields and methods are promoted to the outer struct.

type Person struct {
	Name  string
	Email string
}

func (p Person) Login() {
	fmt.Printf("%s tizimga kirdi\n", p.Name)
}

// admin struct that embeds the Person struct
type Admin struct{
	Person // embedding Person struct
	Level string
}

