package models

import (
	"database/sql"
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
