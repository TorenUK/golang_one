package main

import (
	"log"
	"encoding/json"
	"net/http"
)

// basic HTTP server with a single route.

func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"message": "Hello, world!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}