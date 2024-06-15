package aji

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type TestimonyRequest struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

func CreateTestimony(w http.ResponseWriter, r *http.Request) {

	var req TestimonyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	client, err := genai.NewClient(r.Context(), option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Printf("Failed to create GenAI client: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	aiModel := client.GenerativeModel("gemini-1.0-pro")
	aiResponse, err := aiModel.GenerateContent(r.Context(), genai.Text(req.Content))

	if err != nil {
		log.Printf("Failed to get AI feedback: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// testimony := models.Testimony{
	// 	UserID:     req.UserID,
	// 	Content:    req.Content,
	// 	AIFeedback: aiResponse,
	// 	CreatedAt:  time.Now(),
	// }

	// if err := models.DB.Create(&testimony).Error; err != nil {
	// 	log.Printf("Failed to save testimony: %v", err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// Respond with the created testimony
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
}
