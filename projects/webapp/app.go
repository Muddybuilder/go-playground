package main

import (
	"fmt"
	"net/http"
)

type Greet struct {
	Language string
	Massage  string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	msg := ""
	query := r.URL.Query()
	switch r.Method {
	case http.MethodGet:
		msg = "Get method!"
	case http.MethodPost:
		msg = "Post method!"

	}

	fmt.Fprintf(w, "Hello, there :) "+msg)
	fmt.Fprintf(w, "Query: %s", query.Get("Q"))
	// w.Header().Set("Content-Type", "application/json")
	// greet := Greet{"en", "Hello, there!"}
	// json.NewEncoder(w).Encode(greet)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/Hello", handleRequest)

	http.ListenAndServe(":8000", mux)

}
