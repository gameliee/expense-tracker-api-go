package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
