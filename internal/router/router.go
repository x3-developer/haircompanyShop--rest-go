package router

import (
	"context"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "haircompany-shop-rest/docs"
	"haircompany-shop-rest/internal/modules/v1/category"
	"haircompany-shop-rest/internal/modules/v1/image"
	"haircompany-shop-rest/pkg/database"
	"log"
	"net/http"
	"sync"
)

func NewRouter(appEnv string, db *database.DB, ctx context.Context, wg *sync.WaitGroup) http.Handler {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	image.RegisterV1ImageRoutes(v1)
	category.RegisterV1CategoryRoutes(v1, db, ctx, wg)

	log.Printf("Registering routes for app environment: %s", appEnv)
	if appEnv != "production" {
		mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	}
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	return mux
}
