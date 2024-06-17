package main

import (
	"bebasrokok-be/controllers/aileen"
	"bebasrokok-be/controllers/neo"
	"bebasrokok-be/controllers/ujik"
	"bebasrokok-be/models"
	"log"
	"net/http"
)

func main() {
	err := models.Setup()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/register", neo.Register)
	mux.HandleFunc("/login", neo.Login)
	mux.HandleFunc("/getstreak", ujik.GetStreak)
	mux.HandleFunc("/deletetestimony", ujik.DeleteTestimony)

	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/checkin", aileen.CheckIn)
	protectedMux.HandleFunc("/getcheckindates", aileen.GetCheckInDates)

	mux.Handle("/checkin", neo.JWTMiddleware(protectedMux))
	mux.Handle("/getcheckindates", neo.JWTMiddleware(protectedMux))

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed to start:%v", err)
	}
}
