package main

type User interface {
	do()
}

type Animal interface {
	do()
}

func main() {

	// 	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 	deleteElementNumber := 4
	// 	s := remove(sortedSlice, deleteElementNumber)
	// 	fmt.Printf("The element index %d deleted from slice", deleteElementNumber)
	// 	fmt.Println(s)
	//
	// 	unSortedSlice := []int{12, 3, 123, 4, 564, 87, 90, 2}
	// 	s2 := unsortedRemove(unSortedSlice, deleteElementNumber)
	// 	fmt.Println(s2)

}

// remove elemnt from sorted slice
func remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// remove element from unsorted slice
func unsortedRemove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
