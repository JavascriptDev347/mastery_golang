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
	fmt.Println("Z: ", t.Z)
}
