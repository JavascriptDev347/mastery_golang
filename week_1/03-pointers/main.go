package main

import "log"

func updateUserID(userIDPtr *int, id int) {
	*userIDPtr = id
}

func updateUserIDCopy(userIDPtr int, id int) {
	userIDPtr = id
	log.Println(userIDPtr)
}
func main() {

	var item *int
	*item = 89
	log.Println(*item)

	// userID := 923
	// userIDPtr := &userID
	// log.Println(userIDPtr)
	// updateUserID(&userID, 122)
	// log.Println(*userIDPtr)

	// updateUserIDCopy(userID, 132)
	// log.Println(userID)

	// *anotherTruckID = 90
	// log.Println(truckID)
	// log.Println(*anotherTruckID)

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
