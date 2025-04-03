package main

import (
	"log"
	"net/http"
	"time"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s\n", time.Now().Format("15:04:05"), r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
