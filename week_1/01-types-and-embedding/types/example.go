package typesandembedding

import "fmt"

type Vector struct {
	X float32
	Y float32
}

type ThreeD struct {
	Z      float32
	Vector //embed
}

func (t ThreeD) Print() {
	fmt.Printf("Three d values are: %v, %v and %v\n", t.X, t.Y, t.Z)
}
