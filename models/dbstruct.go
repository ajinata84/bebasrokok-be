package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Article struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

type EmbedGraph struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	EmbedLink string    `json:"embed_link"`
}

type Testimony struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     int       `json:"user_id"`
	Content    string    `json:"content"`
	AIFeedback string    `json:"ai_feedback"`
}

type Tracker struct {
	TrackerID   int       `json:"tracker_id"`
	UserID      int       `json:"user_id"`
	CheckInDate time.Time `json:"check_in_date"`
}

type User struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Streak   int    `json:"streak"`
}

func CreateUser(db *sql.DB, username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (username, email, password, streak) VALUES (?, ?, ?, 0)"
	_, err = db.Exec(query, username, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	user := &User{}
	query := "SELECT userid, username, email, password FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&user.UserID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func GetAllTestimonies(db *sql.DB) ([]Testimony, error) {
	query := "SELECT id, created_at, userid, content, aiFeedback FROM testimony"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testimonies []Testimony
	for rows.Next() {
		var testimony Testimony
		if err := rows.Scan(&testimony.ID, &testimony.CreatedAt, &testimony.UserID, &testimony.Content, &testimony.AIFeedback); err != nil {
			return nil, err
		}
		testimonies = append(testimonies, testimony)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return testimonies, nil
}
