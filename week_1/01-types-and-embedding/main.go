package main

import (
	"fmt"

	types "github.com/JavascriptDev347/mastery_golang/week_1/01-types-and-embedding/types"
)

func main() {
	a := types.Admin{
		Person: types.Person{
			Name:  "John",
			Email: "john@gmail.com",
		},
		Level: "1",
	}

	fmt.Println("Hello: ", a.Name)
}
