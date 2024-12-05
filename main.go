package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Extracted handler
func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Correct header key
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode("hello this is a very dummy project"); err != nil {
		http.Error(w, "Something Went wrong", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", MainHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Something went wrong while starting the server %v", err)
	}
}
