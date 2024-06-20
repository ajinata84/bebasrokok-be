package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Setup() error {
	dsn := "root:KontolKuda@tcp(172.29.145.89:3306)/rokok?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error verifying connection to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database")
	return nil
}

// GetDB returns a reference to the database
func GetDB() *sql.DB {
	return db
}

// GetUserByID fetches a user by their ID
func GetUserByID(userID int) (*User, error) {
	var user User
	query := "SELECT user_id, username, email, password FROM users WHERE user_id = ?"
	row := db.QueryRow(query, userID)

	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// DeleteTestimonyByID deletes a testimony by its ID
func DeleteTestimonyByID(testimonyID int) error {
	query := "DELETE FROM testimony WHERE id = ?"
	result, err := db.Exec(query, testimonyID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("KONTOL")
	}

	return nil
}
