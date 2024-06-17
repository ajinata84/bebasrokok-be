package aji

import (
	"bebasrokok-be/controllers/neo"
	"bebasrokok-be/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var jwtKey = []byte("KUDA")

type TestimonyRequest struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

type TestimonyResponse struct {
	UserID     int    `json:"user_id"`
	Content    string `json:"content"`
	AiResponse any    `json:"AiResponse"`
}

func CreateTestimony(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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

	var req TestimonyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	testimony := "berikan tanggapan dari testimoni untuk berhenti merokok dibawah\n" + req.Content
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(""))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.0-pro")
	cs := model.StartChat()

	send := func(msg string) *genai.GenerateContentResponse {
		fmt.Printf("== Me: %s\n== Model:\n", msg)
		res, err := cs.SendMessage(ctx, genai.Text(msg))
		if err != nil {
			log.Fatal(err)
		}
		return res
	}

	aiRes := send(testimony)

	db := models.GetDB()

	query := "INSERT INTO testimony (userid, content, aiFeedback) VALUES (?, ?, ?)"

	_, err = db.Exec(query, userID, req.Content, aiRes.Candidates[0].Content.Parts[0])

	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting testimony data: %v", err), http.StatusInternalServerError)
		return
	}

	response := TestimonyResponse{
		UserID:     userID,
		Content:    req.Content,
		AiResponse: aiRes.Candidates[0].Content.Parts[0],
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func EditTestimony(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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

	var req TestimonyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	testimony := "berikan tanggapan dari testimoni untuk berhenti merokok dibawah\n" + req.Content
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(""))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.0-pro")
	cs := model.StartChat()

	send := func(msg string) *genai.GenerateContentResponse {
		fmt.Printf("== Me: %s\n== Model:\n", msg)
		res, err := cs.SendMessage(ctx, genai.Text(msg))
		if err != nil {
			log.Fatal(err)
		}
		return res
	}

	aiRes := send(testimony)

	db := models.GetDB()

	query := "UPDATE testimony SET content = ?, aiFeedback = ? WHERE userid = ?"

	_, err = db.Exec(query, req.Content, aiRes.Candidates[0].Content.Parts[0], userID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating testimony data: %v", err), http.StatusInternalServerError)
		return
	}

	response := TestimonyResponse{
		UserID:     userID,
		Content:    req.Content,
		AiResponse: aiRes.Candidates[0].Content.Parts[0],
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func TestAI() {
	testimony := "Setelah mengalami serangan jantung ringan, saya menyadari betapa bahayanya merokok bagi kesehatan saya. Dokter saya memberikan beberapa opsi untuk membantu saya berhenti, dan saya memilih kombinasi obat-obatan dan konseling. Hari-hari pertama sangat berat, dan saya mengalami banyak godaan untuk kembali merokok. Tetapi setiap kali saya merasa ingin merokok, saya ingatkan diri saya akan alasan kesehatan saya dan keluarga saya. Sekarang, setelah tiga tahun bebas rokok, saya merasa jauh lebih sehat dan lebih kuat. Berhenti merokok tidak hanya menyelamatkan hidup saya, tetapi juga memberikan saya kesempatan untuk menikmati hidup lebih lama dengan orang-orang yang saya cintai."
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyDCLHX3KqC-zmSv3N2smyc-hfCPIVMlySQ"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.0-pro")
	cs := model.StartChat()

	send := func(msg string) *genai.GenerateContentResponse {
		fmt.Printf("== Me: %s\n== Model:\n", msg)
		res, err := cs.SendMessage(ctx, genai.Text(msg))
		if err != nil {
			log.Fatal(err)
		}
		return res
	}

	res := send(testimony)
	printResponse(res)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
