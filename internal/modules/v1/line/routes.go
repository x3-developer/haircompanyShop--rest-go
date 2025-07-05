package line

import (
	"haircompany-shop-rest/internal/container"
	"haircompany-shop-rest/internal/middleware"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1LineRoutes(mux *http.ServeMux, container *container.Container) {
	repo := NewRepository(container.DB)
	svc := NewService(repo)
	h := NewHandler(svc)

	mux.Handle("/line/create",
		middleware.ChainMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodPost:
					h.Create(w, r)
				default:
					msg := "Method not allowed. Allowed methods: POST"
					response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
				}
			}),
			middleware.DashboardRoleMiddleware("admin"),
			middleware.DashboardAuthMiddleware(container.JWTService),
		),
	)

	mux.HandleFunc("/line", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w)
		default:
			msg := "Method not allowed. Allowed methods: GET"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})

	mux.HandleFunc("/line/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetById(w, r)
		default:
			msg := "Method not allowed. Allowed methods: GET"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})

	mux.Handle("/line/{id}/update",
		middleware.ChainMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodPatch:
					h.Update(w, r)
				default:
					msg := "Method not allowed. Allowed methods: PATCH"
					response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
				}
			}),
			middleware.DashboardRoleMiddleware("admin"),
			middleware.DashboardAuthMiddleware(container.JWTService),
		),
	)

	mux.Handle("/line/{id}/delete",
		middleware.ChainMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodDelete:
					h.Delete(w, r)
				default:
					msg := "Method not allowed. Allowed methods: PATCH"
					response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
				}
			}),
			middleware.DashboardRoleMiddleware("admin"),
			middleware.DashboardAuthMiddleware(container.JWTService),
		),
	)
}
