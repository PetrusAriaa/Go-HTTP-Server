package api

import (
	"github.com/PetrusAriaa/go-http-server/controller"

	"github.com/gorilla/mux"
)

func BookRouter(prefix string, r *mux.Router) {
	b := r.PathPrefix(prefix).Subrouter()
	b.HandleFunc("/", controller.GetBook).Methods("GET")
	b.HandleFunc("/{id}", controller.GetBookById).Methods("GET")
	b.HandleFunc("/", controller.AddBook).Methods("POST")
	b.HandleFunc("/{id}", controller.UpdateBook).Methods("PUT")
	b.HandleFunc("/{id}", controller.DeleteBook).Methods("DELETE")
}
