package image

import (
	"haircompany-shop-rest/internal/container"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func RegisterV1ImageRoutes(mux *http.ServeMux, container *container.Container) {
	svc := NewService(container.FileService)
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
