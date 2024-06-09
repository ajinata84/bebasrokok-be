package main

import (
	"bebasrokok-be/controllers/aileen"
	"bebasrokok-be/models"
	"log"
	"net/http"
)

func main() {
	// Setup database
	err := models.Setup()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Define routes
	// Example route:
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/checkin", aileen.CheckIn)

	// Start the HTTP server
	port := ":8080"
	log.Printf("Starting server on port %s", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
