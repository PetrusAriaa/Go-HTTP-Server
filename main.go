package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PetrusAriaa/go-http-server/api"
	"github.com/gorilla/mux"
)

func Main() {
	r := mux.NewRouter()
	api.BookRouter("/api/books", r)
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	fmt.Println("server listening on port 8080")
	log.Fatal(s.ListenAndServe())
}
