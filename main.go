package main

import (
	"log"
	"net/http"

	"github.com/PetrusAriaa/go-http-server/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.BookRouter("/api/books", r)
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())
}
