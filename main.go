package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PetrusAriaa/go-http-server/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.BookRouter("/api/books", r)
	s := &http.Server{
		Addr:    ":80",
		Handler: r,
	}
	fmt.Println("server listening on port 80")
	log.Fatal(s.ListenAndServe())
}
