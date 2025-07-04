package middleware

import (
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func DashboardRoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	roleSet := make(map[string]struct{})
	for _, role := range allowedRoles {
		roleSet[role] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value("dashboardClaims").(*services.DashboardClaims)
			if !ok || claims == nil {
				response.SendError(w, http.StatusForbidden, "Forbidden", response.Forbidden)
				return
			}

			if _, ok := roleSet[claims.Role]; !ok {
				response.SendError(w, http.StatusForbidden, "Access denied", response.Forbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
