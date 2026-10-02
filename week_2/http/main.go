package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Write([]byte("Hello from the server"))
		return
	case "POST":
		w.Write([]byte("POST request received"))
		return
	case "PUT":
		w.Write([]byte("PUT request received"))
		return
	case "DELETE":
		w.Write([]byte("DELETE request received"))
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	s := &server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}
