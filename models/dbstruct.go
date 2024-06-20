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

type FetchTestimony struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     int       `json:"user_id"`
	Content    string    `json:"content"`
	AIFeedback string    `json:"ai_feedback"`
	Username   string    `json:"username"`
	Age        int       `json:"age"`
}

type FetchAllTestimony struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     int       `json:"user_id"`
	Content    string    `json:"content"`
	AIFeedback string    `json:"ai_feedback"`
	Username   string    `json:"username"`
	Age        int       `json:"age"`
	MaxStreak  int       `json:"max_streak"`
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

func GetUserTestimonies(db *sql.DB, userID int) ([]FetchTestimony, error) {
	query := `
		SELECT t.id, t.created_at, t.userid, t.content, t.aiFeedback, u.username, u.age 
		FROM testimony t
		JOIN users u ON t.userid = u.userid
		WHERE t.userid = ?
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testimonies []FetchTestimony
	for rows.Next() {
		var testimony FetchTestimony
		if err := rows.Scan(&testimony.ID, &testimony.CreatedAt, &testimony.UserID, &testimony.Content, &testimony.AIFeedback, &testimony.Username, &testimony.Age); err != nil {
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

// GetCheckInDates retrieves the check-in dates for a given user ID
func GetCheckInDates(db *sql.DB, userID int) ([]time.Time, error) {
	query := "SELECT checkInDate FROM tracker WHERE userid = ? ORDER BY checkInDate ASC"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checkInDates []time.Time
	for rows.Next() {
		var checkInDate time.Time
		if err := rows.Scan(&checkInDate); err != nil {
			return nil, err
		}
		checkInDates = append(checkInDates, checkInDate)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return checkInDates, nil
}

// CalculateStreak calculates the longest streak of consecutive check-in dates
func CalculateStreak(checkInDates []time.Time) int {
	if len(checkInDates) == 0 {
		return 0
	}

	streak := 1
	maxStreak := 1

	for i := 1; i < len(checkInDates); i++ {
		if checkInDates[i-1].Add(24 * time.Hour).Equal(checkInDates[i]) {
			streak++
		} else {
			if streak > maxStreak {
				maxStreak = streak
			}
			streak = 1
		}
	}

	if streak > maxStreak {
		maxStreak = streak
	}

	return maxStreak
}

// GetUsernameByID retrieves the username for a given user ID
func GetUsernameByID(db *sql.DB, userID int) (string, error) {
	var username string
	query := "SELECT username FROM users WHERE userid = ?"
	err := db.QueryRow(query, userID).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func GetAllTestimonies(db *sql.DB) ([]FetchAllTestimony, error) {
	query := `
		SELECT t.id, t.created_at, t.userid, t.content, t.aiFeedback, u.username, u.age 
		FROM testimony t
		JOIN users u ON t.userid = u.userid
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testimonies []FetchAllTestimony
	for rows.Next() {
		var testimony FetchAllTestimony
		if err := rows.Scan(&testimony.ID, &testimony.CreatedAt, &testimony.UserID, &testimony.Content, &testimony.AIFeedback, &testimony.Username, &testimony.Age); err != nil {
			return nil, err
		}

		// Get check-in dates for the user
		checkInDates, err := GetCheckInDates(db, testimony.UserID)
		if err != nil {
			return nil, err
		}

		// Calculate the maximum streak for the user
		maxStreak := CalculateStreak(checkInDates)
		testimony.MaxStreak = maxStreak

		testimonies = append(testimonies, testimony)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return testimonies, nil
}
