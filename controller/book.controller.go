package controller

import (
	"net/http"

	"encoding/json"

	"github.com/PetrusAriaa/go-http-server/models"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookData := &[]models.Book{
		{Title: "Laut", Author: "Leila Salikha Chudori", Page: 379, Rate: 97},
		{Title: "Dilan: Dia Adalah Dilanku Tahun 1990", Author: "Pidi Baiq", Page: 332, Rate: 87},
		{Title: "Cantik Itu Luka", Author: "Eka Kurniawan", Page: 552, Rate: 93},
		{Title: "Bumi Manusia", Author: "Pramoedya Ananta Toer", Page: 418, Rate: 94},
		{Title: "Bumi", Author: "Tere Liye", Page: 440, Rate: 98},
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(bookData)
}
