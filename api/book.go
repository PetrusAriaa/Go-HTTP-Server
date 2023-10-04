package api

import (
	"github.com/PetrusAriaa/Go-HTTP-Server/controller"

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
