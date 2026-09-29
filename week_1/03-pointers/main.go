package main

import "log"

func main() {

	truckID := 42
	anotherTruckID := &truckID
	log.Println(&truckID)
	log.Println(anotherTruckID)

	*anotherTruckID = 90
	log.Println(truckID)
	log.Println(*anotherTruckID)

	// log.Println(&anotherTruckID)
	// log.Println(truckID)
	// log.Println(*anotherTruckID)
	// truckID = 90
	// log.Println(truckID)
	// log.Println(*anotherTruckID)
}

// a := makeMultiplier(3)
// for i := 1; i <= 9; i++ {
// 	fmt.Println(a(i))
// }
//
// func makeMultiplier(factor int) func(int) int {
// 	countOfCalls := 0

// 	return func(i int) int {
// 		countOfCalls += 1
// 		bonus := 0
// 		if countOfCalls%3 == 0 {
// 			bonus = 100
// 		}

// 		return factor*i + bonus
// 	}
// }
