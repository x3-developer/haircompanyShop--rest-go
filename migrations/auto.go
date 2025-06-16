package main

import (
	"github.com/joho/godotenv"
	"haircompany-shop-rest/config"
	"haircompany-shop-rest/internal/modules/v1/category"
	"haircompany-shop-rest/pkg/database"
	"log"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		if err := godotenv.Load(".env.production.local"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg := config.LoadConfig()
	db := database.NewDB(cfg)

	if err := db.AutoMigrate(&category.Category{}); err != nil {
		log.Fatalf("Error auto migrating: %s", err)
	}
}
