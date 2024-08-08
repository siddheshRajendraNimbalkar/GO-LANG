package main

import (
	"log"
	"net/http"
	"time"
)

func middlewareRoutes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Method: %s, Duration: %v\n", r.Method, duration)
	})
}

func main() {
	route := http.NewServeMux()

	route.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	Server := http.Server{
		Handler: middlewareRoutes(route),
		Addr:    ":4000",
	}
	log.Println("Server on", Server.Addr)
	Server.ListenAndServe()
}
