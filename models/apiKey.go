package models

import "os"

var apiKey = os.Getenv("GEMINI_API_KEY")

func GetApiKey() string {
	return apiKey
}
