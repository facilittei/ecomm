package main

import (
	"fmt"
	"github.com/facilittei/ecomm/internal/config"
	"github.com/facilittei/ecomm/internal/routes"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()
	routes := routes.NewApi()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: routes.Expose(),
	}

	log.Printf("Listening on port %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start routes %v", err)
	}
}
