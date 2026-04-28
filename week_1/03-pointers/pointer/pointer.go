package pointer

import "fmt"

func main() {

}

func functionPointer(){}

func slicePointer() {
	var slice = []int{1, 2, 3, 4, 5}
	var sliceCopy = slice

	sliceCopy[0] = 100

	fmt.Printf("The value of slice: %v\n", slice)         // 100, 2, 3, 4, 5
	fmt.Printf("The value of sliceCopy: %v\n", sliceCopy) // 100, 2, 3, 4, 5

	slice1 := []int{1, 2, 3, 4, 5}
	slice1Copy := make([]int, len(slice1))
	copy(slice1Copy, slice1)

	slice1Copy[0] = 100

	fmt.Printf("The value of slice1: %v\n", slice1)         // 1, 2, 3, 4, 5
	fmt.Printf("The value of slice1Copy: %v\n", slice1Copy) // 100, 2, 3, 4, 5

}

func pointer() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value of nil pointer: %v\n", *p)
	fmt.Printf("The value of int32: %v\n", i)

	p = &i
	i = 90
	fmt.Printf("The value of pointer: %v\n", p)
	fmt.Printf("The value of pointer dereference: %v\n", *p)
}
