package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger middleware for http requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

		next.ServeHTTP(w, r)

		log.Printf("Completed %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	})
}
