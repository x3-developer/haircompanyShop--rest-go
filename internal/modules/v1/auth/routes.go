package auth

import (
	"haircompany-shop-rest/internal/container"
	"haircompany-shop-rest/internal/modules/v1/client_user"
	"haircompany-shop-rest/internal/modules/v1/dashboard_user"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1AuthRoutes(mux *http.ServeMux, container *container.Container) {
	dashboardUserRepo := dashboard_user.NewRepository(container.DB)
	clientUserRepo := client_user.NewRepository(container.DB)
	svc := NewService(container.RedisService, container.JWTService, container.PasswordService, dashboardUserRepo, clientUserRepo)
	h := NewHandler(svc)

	mux.HandleFunc("/auth/dashboard/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.DashboardLogin(w, r)
		default:
			msg := "Method not allowed. Allowed methods: POST"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})

	mux.HandleFunc("/auth/dashboard/refresh", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.DashboardRefreshToken(w, r)
		default:
			msg := "Method not allowed. Allowed methods: POST"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})
}
