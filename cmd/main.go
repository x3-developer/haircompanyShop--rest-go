package main

import (
	"context"
	"github.com/joho/godotenv"
	"haircompany-shop-rest/config"
	"haircompany-shop-rest/internal/middleware"
	"haircompany-shop-rest/internal/router"
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/database"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	loadEnv()
	cfg := config.LoadConfig()
	db := database.NewDB(cfg)

	srv := newHTTPServer(cfg, db, ctx, &wg)
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

func newHTTPServer(cfg *config.Config, db *database.DB, ctx context.Context, wg *sync.WaitGroup) *http.Server {
	r := router.NewRouter(cfg, db, ctx, wg)
	r = middleware.ChainMiddleware(r, middleware.LoggingMiddleware, middleware.RecoverMiddleware, middleware.CORSMiddleware(cfg.CORS), middleware.APIMiddleware(cfg.AuthAppKey))

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
