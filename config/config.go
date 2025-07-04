package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppEnv          string
	AppPort         string
	DbHost          string
	DbPort          string
	DbName          string
	DbUser          string
	DbPassword      string
	DbSsl           string
	CORS            string
	AuthAppKey      string
	DashboardSecret string
	ClientSecret    string
	RedisAddr       string
	RedisPassword   string
	RedisDB         int
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

	authAppKey := os.Getenv("AUTH_APP_KEY")
	if authAppKey == "" {
		log.Fatal("AUTH_APP_KEY environment isn't set")
	}

	dashboardSecret := os.Getenv("JWT_DASHBOARD_SECRET_KEY")
	if dashboardSecret == "" {
		log.Fatal("JWT_DASHBOARD_SECRET_KEY environment isn't set")
	}

	clientSecret := os.Getenv("JWT_CLIENT_SECRET_KEY")
	if clientSecret == "" {
		log.Fatal("JWT_CLIENT_SECRET_KEY environment isn't set")
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatal("REDIS_ADDR environment isn't set")
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisDB := os.Getenv("REDIS_DB")
	if redisDB == "" {
		redisDB = "0"
	}
	redisDBInt, err := strconv.Atoi(redisDB)
	if err != nil {
		log.Fatal("Invalid REDIS_DB value: ", err)
	}

	return &Config{
		AppEnv:          appEnv,
		AppPort:         appPort,
		DbHost:          dbHost,
		DbPort:          dbPort,
		DbName:          dbName,
		DbUser:          dbUser,
		DbPassword:      dbPassword,
		DbSsl:           dbSsl,
		CORS:            corsAllowedOrigins,
		AuthAppKey:      authAppKey,
		DashboardSecret: dashboardSecret,
		ClientSecret:    clientSecret,
		RedisAddr:       redisAddr,
		RedisPassword:   redisPassword,
		RedisDB:         redisDBInt,
	}
}
