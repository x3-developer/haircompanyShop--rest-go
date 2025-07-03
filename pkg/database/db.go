package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"haircompany-shop-rest/config"
	"log"
)

type DB struct {
	*gorm.DB
}

func NewDB(cfg *config.Config) *DB {
	dsn := GetDSN(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	return &DB{db}
}

func GetDSN(cfg *config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbSsl,
	)
}
