package middleware

import (
	"haircompany-shop-rest/pkg/response"
	"net/http"
)

func APIMiddleware(authAppKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-AUTH-APP") != authAppKey {
				msg := "unauthorized request"
				response.SendError(w, http.StatusForbidden, msg, response.Forbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
