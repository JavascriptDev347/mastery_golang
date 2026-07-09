package main

import "fmt"

func main() {
	fmt.Println("Channels")
	ch := make(chan int)
	go func() {
		ch <- 100
	}()

	result := <-ch
	fmt.Println(result)

}
