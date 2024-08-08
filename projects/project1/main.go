package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	staticFile := http.FileServer(http.Dir("./static"))

	router.Handle("/", staticFile)

	router.HandleFunc("POST /form", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "err %v", err)
			return
		}
		name := r.FormValue("name")
		password := r.FormValue("password")
		fmt.Println("name:", name, "password:", password)
		fmt.Fprintf(w, "Received form data: name=%s, password=%s", name, password)
	})

	Server := http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	log.Println("Server on", Server.Addr)
	Server.ListenAndServe()
}
