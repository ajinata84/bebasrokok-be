package aileen

import (
	"bebasrokok-be/controllers/neo"
	"bebasrokok-be/models"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("KUDA")

func CheckIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	claims := &neo.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID

	db := models.GetDB()
	query := "INSERT INTO tracker (userid, checkInDate) VALUES (?, ?)"
	_, err = db.Exec(query, userID, time.Now())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting check-in data: %v", err), http.StatusInternalServerError)
		return
	}

	// Send a successful response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Check-in successful"))
}
