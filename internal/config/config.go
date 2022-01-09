package config

import (
	"os"
	"strconv"
)

// Config holds application settings
type Config struct {
	Port           string
	SqlDsn         string
	SqlMaxOpenConn int
	SqlMaxIdleConn int
	SqlMaxIdleTime string
	RedisHost      string
	RedisPort      string
}

// NewConfig creates an instance of Config
func NewConfig() Config {
	var sqlMaxOpenConn, sqlMaxIdleConn int
	var err error

	sqlMaxIdleConn, err = strconv.Atoi(os.Getenv("SQL_MAX_OPEN_CONN"))
	if err != nil {
		sqlMaxOpenConn = 20
	}

	sqlMaxIdleConn, err = strconv.Atoi(os.Getenv("SQL_MAX_IDLE_CONN"))
	if err != nil {
		sqlMaxIdleConn = 20
	}

	return Config{
		Port:           os.Getenv("PORT"),
		SqlDsn:         os.Getenv("SQL_DSN"),
		SqlMaxOpenConn: sqlMaxOpenConn,
		SqlMaxIdleConn: sqlMaxIdleConn,
		SqlMaxIdleTime: os.Getenv("SQL_MAX_IDLE_TIME"),
		RedisHost:      os.Getenv("REDIS_HOST"),
		RedisPort:      os.Getenv("REDIS_PORT"),
	}
}
