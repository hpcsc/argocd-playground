package middlewares

import (
	"context"
	"net/http"
	"os"
)

const (
	ContextDataKey = "data"
)

func WithHttpContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextDataKey, map[string]string{
			"env": os.Getenv("ENVIRONMENT"),
		})
		requestWithContext := r.WithContext(ctx)

		next.ServeHTTP(w, requestWithContext)
	})
}
