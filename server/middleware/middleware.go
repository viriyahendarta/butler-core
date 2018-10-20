package middleware

import "net/http"

//Middleware holds contract for all middleware implementation
type Middleware interface {
	Middleware(next http.Handler) http.Handler
}
