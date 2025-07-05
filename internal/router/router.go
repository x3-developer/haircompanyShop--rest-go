package router

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"haircompany-shop-rest/config"
	_ "haircompany-shop-rest/docs"
	"haircompany-shop-rest/internal/container"
	"haircompany-shop-rest/internal/middleware"
	"haircompany-shop-rest/internal/modules/v1/auth"
	"haircompany-shop-rest/internal/modules/v1/category"
	"haircompany-shop-rest/internal/modules/v1/dashboard_user"
	"haircompany-shop-rest/internal/modules/v1/image"
	"haircompany-shop-rest/internal/modules/v1/line"
	"haircompany-shop-rest/internal/modules/v1/product_type"
	"net/http"
)

func NewRouter(cfg *config.Config, container *container.Container) http.Handler {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	auth.RegisterV1AuthRoutes(v1, container)
	image.RegisterV1ImageRoutes(v1, container)
	category.RegisterV1CategoryRoutes(v1, container)
	dashboard_user.RegisterV1DashboardUserRoutes(v1, container)
	line.RegisterV1LineRoutes(v1, container)
	product_type.RegisterV1ProductTypeRoutes(v1, container)

	apiHandler := middleware.ChainMiddleware(
		v1,
		middleware.APIMiddleware(cfg.AuthAppKey),
		middleware.CORSMiddleware(cfg.CORS),
	)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiHandler))

	if cfg.AppEnv != "production" {
		mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	}

	return mux
}
