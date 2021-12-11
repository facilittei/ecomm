package config

import "os"

// Config holds application settings
type Config struct {
	Port      string
	RedisHost string
	RedisPort string
}

// NewConfig creates an instance of Config
func NewConfig() Config {
	return Config{
		Port:      os.Getenv("PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}
}
