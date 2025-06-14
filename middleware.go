package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "isLogged", true)
		r = r.WithContext(ctx)
		log.Print(time.Now(), r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-AUTH-TOKEN") != "secretlifeofpets" {
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "isAuthenticated", true)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
