package main

import (
	"fmt"
	"github.com/facilittei/ecomm/internal/config"
	"github.com/facilittei/ecomm/internal/logging"
	"github.com/facilittei/ecomm/internal/routes"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	logger := logging.NewZeroLogger()
	routes := routes.NewApi(logger, rdb)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: routes.Expose(),
	}

	log.Printf("Listening on port %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start routes %v", err)
	}
}
