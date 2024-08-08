package main

import (
	"log"
	"net/http"
)

func main() {
	route := http.NewServeMux()

	route.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi There"))
	})

	route.HandleFunc("GET /user/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("HELLO THERE!!!!")
	})

	route.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/user/" + r.PathValue("id")))
	})

	Server := http.Server{
		Addr:    ":4000",
		Handler: route,
	}

	log.Println("Server is Listening on ", Server.Addr)
	Server.ListenAndServe()
}
