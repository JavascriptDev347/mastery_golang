package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello world ")
		fmt.Println("R:", r.Method)
		fmt.Println("W:", w)
	})

	http.ListenAndServe(":8080", nil)
}
