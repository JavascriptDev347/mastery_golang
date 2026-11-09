package main

import "fmt"

type User interface {
	do()
}

type Animal interface {
	do()
}

func main() {

	sonlar := []int{1, 2, 3, 4, 5}
	for chap, ong := 0, len(sonlar)-1; chap < ong; chap, ong = chap+1, ong-1 {
		// chap = 0, o'ng = len(sonlar)-1 = 4, chap = chap(0+1), ong=
		//
		//
		//
		//
		sonlar[chap], sonlar[ong] = sonlar[ong], sonlar[chap]
	}
	fmt.Println(sonlar)

	// 	var bosh []int
	// 	mevalar := []string{"olma", "anor", "uzum"}
	// 	sonlar := make([]int, 3, 5)
	//
	// 	fmt.Println(bosh, len(bosh), cap(bosh), bosh == nil)
	// 	fmt.Println(mevalar)
	// 	fmt.Println(sonlar, len(sonlar), cap(sonlar))

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
