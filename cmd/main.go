package main

import (
	"context"
	"github.com/joho/godotenv"
	"haircompany-shop-rest/config"
	"haircompany-shop-rest/internal/container"
	"haircompany-shop-rest/internal/middleware"
	"haircompany-shop-rest/internal/router"
	"haircompany-shop-rest/internal/services"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// @title						Hair Company Shop API
// @version					1.0
// @description				REST API для магазина бренда Hair Company
// @contact.name				API Support
// @contact.email				x3.na.tri@gmail.com
// @schemes					http https
// @securityDefinitions.apiKey	BearerAuth
// @in							header
// @name						Authorization
// @description				Введите токен в формате: Bearer {token}
// @securityDefinitions.apiKey	AppAuth
// @in							header
// @name						X-AUTH-APP
// @description				Ключ аутентификации приложения из AUTH_APP_KEY
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	loadEnv()
	cfg := config.LoadConfig()
	diContainer := container.NewContainer(cfg, ctx, &wg)

	srv := newHTTPServer(cfg, diContainer)
	runServer(srv)

	scheduler := services.NewScheduler(ctx, &wg)
	runSchedule(scheduler)

	<-ctx.Done()
	gracefulShutdown(srv, &wg)
}

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		if err := godotenv.Load(".env.production.local"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func newHTTPServer(cfg *config.Config, container *container.Container) *http.Server {
	r := router.NewRouter(cfg, container)
	r = middleware.ChainMiddleware(r, middleware.RecoverMiddleware, middleware.LoggingMiddleware)

	return &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}
}

func runServer(srv *http.Server) {
	go func() {
		log.Printf("Starting server on :%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Stopped listening server: %v", err)
		}
	}()
}

func gracefulShutdown(srv *http.Server, wg *sync.WaitGroup) {
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}

	log.Println("Waiting for background goroutines to finish...")
	wg.Wait()
	log.Println("Server gracefully stopped")
}

func runSchedule(scheduler services.Scheduler) {
	filesystem := services.NewFileSystemService()

	cleanTempTask := scheduler.CreateTask("CleanTempFiles", func() {
		filesystem.CleanTemp()
	})

	scheduler.StartEveryDay(6, 0, cleanTempTask)
}
