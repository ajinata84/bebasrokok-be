package gavin

import (
	"bebasrokok-be/models"
	"encoding/json"
	"net/http"
)

func GetTestimonies(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()

	testimonies, err := models.GetAllTestimonies(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testimonies)
}
