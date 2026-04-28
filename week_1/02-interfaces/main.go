package main

import (
	intern "github.com/JavascriptDev347/mastery_golang/week_1/02-interfaces/inter"
)

func main() {
	c := intern.Circle{
		Radius: 3.12,
	}

	r := intern.Rectangle{
		Width:  3,
		Height: 5,
	}

	intern.PrintArea(c)
	intern.PrintArea(r)
}
