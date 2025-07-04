package dashboard_user

import (
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/database"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1DashboardUserRoutes(mux *http.ServeMux, db *database.DB) {
	repo := NewRepository(db)
	passwordSvc := services.NewPasswordService()
	svc := NewService(repo, passwordSvc)
	h := NewHandler(svc)

	mux.HandleFunc("/dashboard-user/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.Create(w, r)
		default:
			msg := "Method not allowed. Allowed methods: POST"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})
}
