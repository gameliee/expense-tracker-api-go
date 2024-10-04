package main

import (
	"database/sql"
	"log"
)

func InitializeTables(db *sql.DB) error {
	// Create users table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Printf("Error creating users table: %v", err)
		return err
	}

	// Create expenses table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS expenses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			name TEXT NOT NULL,
			description TEXT,
			amount REAL NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		log.Printf("Error creating expenses table: %v", err)
		return err
	}

	log.Println("Database tables initialized successfully")
	return nil
}
