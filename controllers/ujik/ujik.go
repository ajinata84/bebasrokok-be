package ujik

import (
	"bebasrokok-be/models"
	"net/http"
	"strconv"
)

// GetStreak handles the request to get the user's streak
// func GetStreak(w http.ResponseWriter, r *http.Request) {
// 	userIDStr := r.URL.Query().Get("user_id")
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := models.GetUserByID(userID)
// 	if err != nil {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]int{"streak": user.Streak})
// }

// DeleteTestimony handles the request to delete a testimony
func DeleteTestimony(w http.ResponseWriter, r *http.Request) {
	testimonyIDStr := r.URL.Query().Get("id")
	testimonyID, err := strconv.Atoi(testimonyIDStr)
	if err != nil {
		http.Error(w, "Invalid testimony ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteTestimonyByID(testimonyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Testimony deleted successfully"))
}
