package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gamelieelearn/expense-tracker-api-go/config"
)

// InitDB initializes the database connection
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}
