package main

import (
	"database/sql"
	"fmt"
	"log"

	"gamelieelearn/expense-tracker-api-go/config"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the database connection
func InitDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}
