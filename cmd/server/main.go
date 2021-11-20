package main

import (
	"github.com/facilittei/ecomm/internal/config"
	"github.com/facilittei/ecomm/internal/controllers"
	"log"
)

func main() {
	cfg := config.NewConfig()
	app := controllers.NewApp(cfg)

	if err := app.Listen(); err != nil {
		log.Fatalf("could not start server %v", err)
	}
}
