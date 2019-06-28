package globalmiddlewares

import (
	"github.com/tonyooi/go-base-uservice/src/middlewarecontext"
	"context"
	"fmt"
	"net/http"
)

// Authenticate - globalmiddleware for authentication
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var parentContext, newContext context.Context
		fmt.Println("Authentication Successful")
		parentContext = r.Context()
		newContext = middlewarecontext.NewContextWithValue(parentContext, "UserInfo", "Brown")
		r = r.WithContext(newContext)
		next.ServeHTTP(w, r)
	})
}
