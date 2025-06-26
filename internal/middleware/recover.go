package middleware

import (
	"haircompany-shop-rest/pkg/response"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC recovered: %v\n%s", err, debug.Stack())

				msg := "panic occurred while processing the request"
				response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
