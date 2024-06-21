package controllers

import (
	"bebasrokok-be/controllers/neo"
	"bebasrokok-be/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("KUDA")

func GetStreak(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = tokenString[len("Bearer "):]

		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		claims := &neo.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userID := claims.UserID

		if err != nil {
			http.Error(w, "invalid user_id", http.StatusBadRequest)
			return
		}

		checkInDates, err := models.GetCheckInDates(db, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		username, err := models.GetUsernameByID(db, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		streakCount := models.CalculateStreak(checkInDates)

		response := struct {
			Username     string      `json:"username"`
			StreakCount  int         `json:"streak_count"`
			CheckInDates []time.Time `json:"check_in_dates"`
		}{
			Username:     username,
			StreakCount:  streakCount,
			CheckInDates: checkInDates,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
