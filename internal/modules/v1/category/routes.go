package category

import (
	"context"
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/database"
	"haircompany-shop-rest/pkg/response"
	"net/http"
	"sync"
)

func RegisterV1CategoryRoutes(mux *http.ServeMux, db *database.DB, ctx context.Context, wg *sync.WaitGroup) {
	fileSvc := services.NewFileSystemService()
	repo := NewRepository(db)
	svc := NewService(repo, fileSvc, ctx, wg)
	h := NewHandler(svc)

	mux.HandleFunc("/category/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.Create(w, r)
		default:
			msg := "Method not allowed. Allowed methods: POST"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})

	mux.HandleFunc("/category", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w)
		default:
			msg := "Method not allowed. Allowed methods: GET"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})

	mux.HandleFunc("/category/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetById(w, r)
		default:
			msg := "Method not allowed. Allowed methods: GET"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})
}
