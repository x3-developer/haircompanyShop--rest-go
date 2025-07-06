package desired_result

import (
	"haircompany-shop-rest/internal/container"
	"haircompany-shop-rest/internal/middleware"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1DesiredResultRoutes(mux *http.ServeMux, container *container.Container) {
	repo := NewRepository(container.DB)
	svc := NewService(repo)
	h := NewHandler(svc)

	mux.Handle("/desired-result/create",
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

	mux.HandleFunc("/desired-result", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w)
		default:
			msg := "Method not allowed. Allowed methods: GET"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})

	mux.HandleFunc("/desired-result/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetById(w, r)
		default:
			msg := "Method not allowed. Allowed methods: GET"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})

	mux.Handle("/desired-result/{id}/update",
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

	mux.Handle("/desired-result/{id}/delete",
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
