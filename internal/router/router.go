package router

import (
	"haircompany-shop-rest/internal/modules/v1/category"
	"haircompany-shop-rest/pkg/database"
	"net/http"
)

func NewRouter(db *database.DB) *http.ServeMux {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	category.RegisterV1CategoryRoutes(v1, db)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	return mux
}
