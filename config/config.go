package config

type Config struct {
	// Add configuration fields
	DatabasePath string
}

func NewConfig() *Config {
	return &Config{
		DatabasePath: "db.sqlite3",
	}
}
