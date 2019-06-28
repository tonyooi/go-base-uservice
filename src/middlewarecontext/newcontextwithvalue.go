package middlewarecontext

import (
	"context"
)

// NewContextWithValue - it actaullu generates and return a new context to replace the old one
func NewContextWithValue(parentContext context.Context, key, value interface{}) context.Context {
	return context.WithValue(parentContext, key, value)
}
