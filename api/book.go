package api

import (
	"github.com/PetrusAriaa/go-http-server/controller"

	"github.com/gorilla/mux"
)

func BookRouter(prefix string, r *mux.Router) {
	b := r.PathPrefix(prefix).Subrouter()
	b.HandleFunc("/", controller.GetBook)
	b.HandleFunc("/{id}", controller.GetBookById)
}
