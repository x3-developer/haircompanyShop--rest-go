package auth

import (
	"haircompany-shop-rest/internal/modules/v1/client_user"
	"haircompany-shop-rest/internal/modules/v1/dashboard_user"
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/database"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1AuthRoutes(mux *http.ServeMux, db *database.DB, dashboardSecret, clientSecret string) {
	jwtSvc := services.NewJWTService(dashboardSecret, clientSecret)
	passwordSvc := services.NewPasswordService()
	dashboardUserRepo := dashboard_user.NewRepository(db)
	clientUserRepo := client_user.NewRepository(db)
	svc := NewService(jwtSvc, passwordSvc, dashboardUserRepo, clientUserRepo)
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

	//mux.HandleFunc("/auth/dashboard/refresh", func(w http.ResponseWriter, r *http.Request) {
	//	switch r.Method {
	//	case http.MethodPost:
	//		h.Refresh(w, r)
	//	default:
	//		msg := "Method not allowed. Allowed methods: POST"
	//		response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
	//	}
	//})
}
