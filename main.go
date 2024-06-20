package main

import (
	"bebasrokok-be/controllers"
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
	mux.HandleFunc("/deletetestimony", ujik.DeleteTestimony)
	mux.HandleFunc("/viewtestimonies", gavin.GetTestimonies)
	mux.HandleFunc("/viewgraphs", aji.GetGraphs)
	mux.HandleFunc("/viewarticles", aji.GetArticles)

	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/checkin", aileen.CheckIn)
	protectedMux.HandleFunc("/getcheckindates", aileen.GetCheckInDates)
	protectedMux.HandleFunc("/create-testimony", aji.CreateTestimony)
	protectedMux.HandleFunc("/edit-testimony", aji.EditTestimony)
	protectedMux.HandleFunc("/getstreak", controllers.GetStreak(models.GetDB()))
	protectedMux.HandleFunc("/getuser-testimonies", aji.GetUserTestimonies)

	// JWT Middleware applied to the protectedMux

	// CORS middleware function
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}

	// Wrap your mux with the CORS middleware
	handler := corsMiddleware(mux)

	mux.Handle("/getstreak", neo.JWTMiddleware(corsMiddleware(protectedMux)))
	mux.Handle("/checkin", neo.JWTMiddleware(corsMiddleware(protectedMux)))
	mux.Handle("/getcheckindates", neo.JWTMiddleware(corsMiddleware(protectedMux)))
	mux.Handle("/create-testimony", neo.JWTMiddleware(corsMiddleware(protectedMux)))
	mux.Handle("/edit-testimony", neo.JWTMiddleware(corsMiddleware(protectedMux)))
	mux.Handle("/getuser-testimonies", neo.JWTMiddleware(corsMiddleware(protectedMux)))

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	err = http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatalf("Server failed to start:%v", err)
	}
}
