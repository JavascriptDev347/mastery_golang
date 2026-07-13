package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/all", getUsersHandler)
	mux.HandleFunc("/create", createUserHandler)
	mux.HandleFunc("/users/{id}", getUserByIDHandler)

	wrappedMux := loggingMiddleware(mux)
	http.ListenAndServe(":8080", wrappedMux)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of all users")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User created")
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // Go 1.22+ feature
	fmt.Fprintf(w, "User ID: %s\n", id)
}

// middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r) // call the wrapped handler
		log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}
