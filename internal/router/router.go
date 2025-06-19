package router

import (
	"haircompany-shop-rest/internal/modules/v1/category"
	"haircompany-shop-rest/internal/modules/v1/image"
	"haircompany-shop-rest/pkg/database"
	"net/http"
)

func NewRouter(db *database.DB) http.Handler {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	image.RegisterV1ImageRoutes(v1)
	category.RegisterV1CategoryRoutes(v1, db)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	return mux
}
