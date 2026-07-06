package main

import (
	"fmt"
	"slices"
)

func sliceOperator() {
	s := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	fmt.Printf("Real case of %s\n", s)
	fmt.Printf("s[1:4] %v\n", string(s[1:4]))
	fmt.Printf("s[:3] %v\n", string(s[:3]))
	fmt.Printf("s[3:] %v\n", string(s[3:]))

	s1 := s[2:]
	s1[1] = 'x'
	fmt.Printf("s1 %v\n", string(s1))
	fmt.Printf("s %v\n", string(s))

	fmt.Println("---- Length and capacity of slices ----")
	fmt.Printf("len(s1) %d\n", len(s1))
	fmt.Printf("cap(s1) %d\n", cap(s1))

	s2 := append(s1, 12)
	fmt.Printf("s2 %v\n", string(s2))
	fmt.Printf("len(s2) %v", len(s2))
	fmt.Printf("cap(s2) %v", cap(s2))

	fmt.Println("---- copy of slices ----")

	s3 := make([]byte, len(s2))
	copy(s3, s2)
	fmt.Printf("s3 %v\n", string(s3))
	fmt.Printf("len(s3) %v", len(s3))
	fmt.Printf("cap(s3) %v", cap(s3))

	a := []int{1, 2, 3, 4, 5}
	a = append(a, 1000)
	fmt.Println("a)", a)
	b := make([]int, len(a))
	b = append(b, a...)
	fmt.Println("b)", b)

	fmt.Println("a == b?", slices.Equal(a, b))

	fmt.Println("---- MAP in go ----")
	nums := []int{1, 2, 3}
	fmt.Println("Nums", nums)

	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = v * 2
	}
	fmt.Println("Result", result)
	fmt.Println("After map nums:", nums)

}
