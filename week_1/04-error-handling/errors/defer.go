package errors

import "fmt"

func deferFunc() {
	fmt.Println("Start")

	defer fmt.Println("Birinchi defer")
	defer fmt.Println("Ikkinchi defer")
	defer fmt.Println("Uchinchi defer")

	fmt.Println("End")
}
