package config

import (
	"os"
)

type Config struct {
	Mode         string
	DatabasePath string
}

func NewConfig() *Config {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "development"
	}

	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "db.sqlite3"
	}

	if mode == "test" {
		dbPath = ":memory:"
	}

	return &Config{
		Mode:         mode,
		DatabasePath: dbPath,
	}
}
