package main

import (
	"bebasrokok-be/controllers/aileen"
	"bebasrokok-be/controllers/aji"
	"bebasrokok-be/controllers/gavin"
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
	protectedMux.HandleFunc("/create-testimony", aji.CreateTestimony)
	protectedMux.HandleFunc("/edit-testimony", aji.EditTestimony)
	mux.HandleFunc("/viewtestimonies", gavin.GetTestimonies)

	mux.Handle("/checkin", neo.JWTMiddleware(protectedMux))
	mux.Handle("/getcheckindates", neo.JWTMiddleware(protectedMux))
	mux.Handle("/create-testimony", neo.JWTMiddleware(protectedMux))
	mux.Handle("/edit-testimony", neo.JWTMiddleware(protectedMux))

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed to start:%v", err)
	}
}
