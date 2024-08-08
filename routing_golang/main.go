package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Bro struct {
	Name   string `json:name`
	RollNo string `json:rollNo`
}

func main() {
	route := http.NewServeMux()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(Bro{Name: "om", RollNo: "18"})
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	route.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(Bro{Name: "om1", RollNo: "18"})
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	route.HandleFunc("/{id}/{new}", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(Bro{Name: "om3", RollNo: "18"})
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	route.HandleFunc("POST /{id}/{new}/u", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.String()))
	})

	Server := http.Server{
		Addr:    ":4000",
		Handler: route,
	}

	log.Println("Port is on 4000")
	Server.ListenAndServe()
}
