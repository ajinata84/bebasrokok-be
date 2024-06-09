package aileen

import (
	"bebasrokok-be/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CheckInRequest represents the expected request body for a check-in
type CheckInRequest struct {
	UserID int `json:"UserID"` // Use the JSON tag to match the incoming JSON field name
}

// CheckIn handles the user check-in
func CheckIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var req CheckInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the request data
	if req.UserID == 0 {
		http.Error(w, "UserID is required", http.StatusBadRequest)
		return
	}

	// Insert check-in data into the database
	db := models.GetDB()
	query := "INSERT INTO tracker (userid, checkInDate) VALUES (?, ?)"
	_, err := db.Exec(query, req.UserID, time.Now())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting check-in data: %v", err), http.StatusInternalServerError)
		return
	}

	// Send a successful response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Check-in successful"))
}
