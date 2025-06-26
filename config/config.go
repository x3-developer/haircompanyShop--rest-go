package config

import (
	"log"
	"os"
)

type Config struct {
	AppEnv     string
	AppPort    string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	DbSsl      string
	CORS       string
}

func LoadConfig() *Config {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		log.Fatal("APP_ENV environment isn't set")
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT environment isn't set")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB_HOST environment isn't set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB_PORT environment isn't set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment isn't set")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB_USER environment isn't set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD environment isn't set")
	}

	dbSsl := os.Getenv("DB_SSL")
	if dbSsl == "" {
		dbSsl = "verify-full"
	}

	corsAllowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if corsAllowedOrigins == "" {
		log.Fatal("CORS_ALLOWED_ORIGINS environment isn't set")
	}

	return &Config{
		AppEnv:     appEnv,
		AppPort:    appPort,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbSsl:      dbSsl,
		CORS:       corsAllowedOrigins,
	}
}
