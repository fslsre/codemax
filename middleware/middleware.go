package middleware

import (
	"log"
	"net/http"
)

// Middleware middleware
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		h.ServeHTTP(w, r)
	})
}
