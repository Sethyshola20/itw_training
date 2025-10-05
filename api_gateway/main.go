package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func createProxy(target string, allowedOrigin string) http.Handler {
	urlTarget, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(urlTarget)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		if cookie := req.Header.Get("Cookie"); cookie != "" {
			req.Header.Set("Cookie", cookie)
		}
		req.Header.Del("Origin")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		proxy.ServeHTTP(w, r)
	})
}

func main() {
	authURL := os.Getenv("AUTH_URL")
	billingURL := os.Getenv("BILLING_URL")
	allowedOrigin := os.Getenv("UI_URL")

	if authURL == "" || billingURL == "" || allowedOrigin == "" {
		log.Fatal("AUTH_URL, BILLING_URL, and UI_URL environment variables must be set")
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	http.Handle("/api/users", createProxy(authURL, allowedOrigin))
	http.Handle("/api/users/", createProxy(authURL, allowedOrigin))

	http.Handle("/api/signup", createProxy(authURL, allowedOrigin))
	http.Handle("/api/signup/", createProxy(authURL, allowedOrigin))

	http.Handle("/api/token", createProxy(authURL, allowedOrigin))
	http.Handle("/api/token/", createProxy(authURL, allowedOrigin))

	http.Handle("/api/token/refresh", createProxy(authURL, allowedOrigin))
	http.Handle("/api/token/refresh/", createProxy(authURL, allowedOrigin))

	http.Handle("/api/billing/subscriptions", createProxy(billingURL, allowedOrigin))
	http.Handle("/api/billing/subscriptions/", createProxy(billingURL, allowedOrigin))

	http.Handle("/api/billing/invoices", createProxy(billingURL, allowedOrigin))
	http.Handle("/api/billing/invoices/", createProxy(billingURL, allowedOrigin))

	log.Println("API Gateway listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
