package middleware

import (
	"net/http"
	"strings"
)

func CORSMiddleware(corsStr string) func(http.Handler) http.Handler {
	corsAllowedOrigins := strings.Split(corsStr, ",")
	for i, origin := range corsAllowedOrigins {
		corsAllowedOrigins[i] = strings.TrimSpace(origin)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" {
				for _, allowedOrigin := range corsAllowedOrigins {
					if allowedOrigin == origin {
						w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
						w.Header().Set("Vary", "Origin")
						break
					}
				}
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
