package middleware

import (
	"context"
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/response"
	"net/http"
	"strings"
)

func DashboardAuthMiddleware(jwtSvc services.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authedHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authedHeader, "Bearer ") {
				response.SendError(w, http.StatusUnauthorized, "Unauthorized", response.Unauthorized)
				return
			}

			token := strings.TrimPrefix(authedHeader, "Bearer ")
			claims, err := jwtSvc.ValidateDashboardToken(token)
			if err != nil || claims == nil {
				response.SendError(w, http.StatusUnauthorized, "Unauthorized", response.Unauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "dashboardClaims", claims)
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		})
	}
}

func ClientAuthMiddleware(jwtSvc services.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authedHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authedHeader, "Bearer ") {
				response.SendError(w, http.StatusUnauthorized, "Unauthorized", response.Unauthorized)
				return
			}

			token := strings.TrimPrefix(authedHeader, "Bearer ")
			claims, err := jwtSvc.ValidateClientToken(token)
			if err != nil || claims == nil {
				response.SendError(w, http.StatusUnauthorized, "Unauthorized", response.Unauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "clientClaims", claims)
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		})
	}
}
