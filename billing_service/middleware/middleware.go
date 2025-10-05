package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
		if allowedOrigins == "" {
			frontendURL := os.Getenv("FRONTEND_URL")
			if frontendURL != "" {
				allowedOrigins = "http://localhost:5173," + frontendURL
			} else {
				allowedOrigins = "http://localhost:5173"
			}
		}

		requestOrigin := r.Header.Get("Origin")
		if requestOrigin != "" && allowedOrigins != "*" {
			origins := strings.Split(allowedOrigins, ",")
			allowed := false
			for _, origin := range origins {
				if strings.TrimSpace(origin) == requestOrigin {
					allowed = true
					w.Header().Set("Access-Control-Allow-Origin", requestOrigin)
					break
				}
			}

			if !allowed {
				w.Header().Set("Access-Control-Allow-Origin", "null")
			}
		} else {
			// For development or when no Origin header is present
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
