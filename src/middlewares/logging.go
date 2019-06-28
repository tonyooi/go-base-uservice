package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tonyooi/go-base-uservice/src/middlewarecontext"
)

// Logging - middleware for authentication
func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var parentContext, newContext context.Context
		fmt.Println("Logging Turned On")
		parentContext = r.Context()
		newContext = middlewarecontext.NewContextWithValue(parentContext, "LogStatus", true)
		r = r.WithContext(newContext)
		next.ServeHTTP(w, r)
	}
}
