package category

import (
	"haircompany-shop-rest/pkg/database"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1CategoryRoutes(mux *http.ServeMux, db *database.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
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
