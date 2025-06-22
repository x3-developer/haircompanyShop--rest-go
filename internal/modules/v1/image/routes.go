package image

import (
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1ImageRoutes(mux *http.ServeMux) {
	fileSvc := services.NewFileSystemService()
	svc := NewService(fileSvc)
	h := NewHandler(svc)

	mux.HandleFunc("/image/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.Upload(w, r)
		default:
			msg := "Method not allowed. Allowed methods: POST"
			response.SendError(w, http.StatusMethodNotAllowed, msg, response.MethodNotAllowed)
		}
	})
}
