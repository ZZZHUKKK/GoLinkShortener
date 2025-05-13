package middleware

import (
	"net/http"
)

type Middlware func(http.Handler) http.Handler

func Chain(middlewares ...Middlware) Middlware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
