package main

import "fmt"

func example() {
	fmt.Println("Hello map")

	students := make(map[string]int)

	students["Ali"] = 70
	students["Vali"] = 90
	students["Russel"] = 95
	students["Bussel"] = 59
	students["Rose"] = 44

	fmt.Println(students)

	for key, value := range students {
		fmt.Println(key, value)
	}

	for key, value := range students {
		if value > 89 {
			fmt.Printf("%v passed from exam: %v. You got a grade of A.\n", key, value)
		} else if value > 79 {
			fmt.Printf("%v passed from exam: %v. You got a grade of B.\n", key, value)
		} else if value > 69 {
			fmt.Printf("%v passed from exam: %v. You got a grade of C.\n", key, value)
		} else if value > 59 {
			fmt.Printf("%v passed from exam: %v. You got a grade of D.\n", key, value)
		} else {
			fmt.Printf("%v failed from exam: %v\n", key, value)
		}
	}

	grade, ok := students["Bobur"]
	if ok {
		fmt.Printf("Grade for Bobur: %v\n", grade)
	} else {
		fmt.Println("Bobur is not in the map")
	}

	fmt.Println("Length of the map: ", len(students))
}
