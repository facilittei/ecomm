package main

import (
	"database/sql"
	"fmt"
	"github.com/facilittei/ecomm/internal/config"
	"github.com/facilittei/ecomm/internal/logging"
	"github.com/facilittei/ecomm/internal/routes"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()

	sqlClient, err := sql.Open("postgres", cfg.SqlDsn)
	if err != nil {
		log.Fatalf("could not connect to SQL database")
	}

	err = sqlClient.Ping()
	if err != nil {
		log.Fatalf("could not ping SQL database connection")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	logger := logging.NewZeroLogger()
	routes := routes.NewApi(sqlClient, redisClient, logger)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: routes.Expose(),
	}

	log.Printf("Listening on port %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start routes %v", err)
	}
}
