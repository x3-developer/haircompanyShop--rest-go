package router

import (
	httpSwagger "github.com/swaggo/http-swagger"
	_ "haircompany-shop-rest/docs"
	"haircompany-shop-rest/internal/container"
	"haircompany-shop-rest/internal/modules/v1/auth"
	"haircompany-shop-rest/internal/modules/v1/category"
	"haircompany-shop-rest/internal/modules/v1/dashboard_user"
	"haircompany-shop-rest/internal/modules/v1/image"
	"log"
	"net/http"
)

func NewRouter(appEnv string, container *container.Container) http.Handler {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	image.RegisterV1ImageRoutes(v1, container)
	category.RegisterV1CategoryRoutes(v1, container)
	dashboard_user.RegisterV1DashboardUserRoutes(v1, container)
	auth.RegisterV1AuthRoutes(v1, container)

	log.Printf("Registering routes for app environment: %s", appEnv)
	if appEnv != "production" {
		mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	}
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	return mux
}
