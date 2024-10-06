package main

import (
	"gamelieelearn/expense-tracker-api-go/domain"
	"log"

	"gorm.io/gorm"
)

func InitializeTables(db *gorm.DB) (err error) {
	err = db.AutoMigrate(domain.User{}, domain.Expense{})
	log.Println("Database tables initialized successfully")
	return
}
