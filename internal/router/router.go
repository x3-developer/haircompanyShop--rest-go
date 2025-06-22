package router

import (
	"context"
	"haircompany-shop-rest/internal/modules/v1/category"
	"haircompany-shop-rest/internal/modules/v1/image"
	"haircompany-shop-rest/pkg/database"
	"net/http"
	"sync"
)

func NewRouter(db *database.DB, ctx context.Context, wg *sync.WaitGroup) http.Handler {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	image.RegisterV1ImageRoutes(v1)
	category.RegisterV1CategoryRoutes(v1, db, ctx, wg)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	return mux
}
