package dashboard_user

import (
	"haircompany-shop-rest/internal/container"
	"haircompany-shop-rest/internal/middleware"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1DashboardUserRoutes(mux *http.ServeMux, container *container.Container) {
	repo := NewRepository(container.DB)
	svc := NewService(repo, container.PasswordService)
	h := NewHandler(svc)

	mux.Handle("/dashboard-user/create",
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
			middleware.DashboardRoleMiddleware("admin", "manager"),
			middleware.DashboardAuthMiddleware(container.JWTService),
		),
	)
}
