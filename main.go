package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PetrusAriaa/Go-HTTP-Server/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.BookRouter("/api/books", r)
	s := &http.Server{
		Addr:    ":80",
		Handler: r,
	}
	fmt.Println("server listening on http://127.0.0.1:80")
	log.Fatal(s.ListenAndServe())
}
