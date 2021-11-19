package config

import "os"

// Config holds application settings
type Config struct {
	Port string
}

// NewConfig creates an instance of Config
func NewConfig() Config {
	port := os.Getenv("PORT")

	return Config{
		Port: port,
	}
}
