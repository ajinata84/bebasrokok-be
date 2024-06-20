package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Article struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	ThumbnailUrl string    `json:"thumbnail"`
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
	Age      int    `json:"age"`
}

func CreateUser(db *sql.DB, username, email, password string, age int) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (username, email, password, age) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, username, email, hashedPassword, age)
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

func GetAllArticles(db *sql.DB) ([]Article, error) {
	query := "SELECT id, created_at, title, content, thumbnailUrl FROM articles"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.CreatedAt, &article.Title, &article.Content, &article.ThumbnailUrl); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func GetAllEmbedGraphs(db *sql.DB) ([]EmbedGraph, error) {
	query := "SELECT id, created_at, embedLink FROM embedGraphs"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var embedGraphs []EmbedGraph
	for rows.Next() {
		var embedGraph EmbedGraph
		if err := rows.Scan(&embedGraph.ID, &embedGraph.CreatedAt, &embedGraph.EmbedLink); err != nil {
			return nil, err
		}
		embedGraphs = append(embedGraphs, embedGraph)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return embedGraphs, nil
}
