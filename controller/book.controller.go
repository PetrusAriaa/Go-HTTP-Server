package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/PetrusAriaa/go-http-server/lib"
	"github.com/PetrusAriaa/go-http-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

type res struct {
	Data     []*models.Book `json:"data"`
	Length   int64          `json:"length"`
	Accessed time.Time      `json:"accessed"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")

	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	var books []*models.Book
	if err = cursor.All(context.Background(), &books); err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res{Data: books, Length: int64(len(books)), Accessed: time.Now()})
}
