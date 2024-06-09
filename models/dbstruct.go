package models

import (
	"time"
)

// Article represents the articles table
type Article struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

// EmbedGraph represents the embedGraphs table
type EmbedGraph struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	EmbedLink string    `json:"embed_link"`
}

// Testimony represents the testimony table
type Testimony struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     int       `json:"user_id"`
	Content    string    `json:"content"`
	AIFeedback string    `json:"ai_feedback"`
}

// Tracker represents the tracker table
type Tracker struct {
	TrackerID   int       `json:"tracker_id"`
	UserID      int       `json:"user_id"`
	CheckInDate time.Time `json:"check_in_date"`
}

// User represents the users table
type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Streak   int    `json:"streak"`
}
