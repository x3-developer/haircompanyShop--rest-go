package main

import (
	"github.com/joho/godotenv"
	"haircompany-shop-rest/config"
	"haircompany-shop-rest/internal/router"
	"haircompany-shop-rest/pkg/database"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		if err := godotenv.Load(".env.production.local"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg := config.LoadConfig()
	db := database.NewDB(cfg)
	r := router.NewRouter(db)

	log.Printf("Starting server on :%s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
