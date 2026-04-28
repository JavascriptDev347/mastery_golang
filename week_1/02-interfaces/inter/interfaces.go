package inter

import (
	"fmt"
	"math"
)

// Shape Go interfeyslari implicit — agar struct barcha methodlarni implement qilsa, u interfeys uchun "tayyor". Hech qanday implements kerak emas.

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("Maydoni: %0.2f\n", s.Area())
}

// Reader and Writer embedding interfaces
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}
