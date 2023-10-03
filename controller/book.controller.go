package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PetrusAriaa/go-http-server/lib"
	"github.com/PetrusAriaa/go-http-server/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookList struct {
	Data     []*models.Book `json:"data"`
	Length   int64          `json:"length"`
	Accessed time.Time      `json:"accessed"`
}

type SingleBook struct {
	Data     *models.Book `json:"data"`
	Length   int64        `json:"length"`
	Accessed time.Time    `json:"accessed"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	var books []*models.Book

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")

	defer conn.Disconnect(context.Background())
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.Background(), &books); err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(BookList{Data: books, Length: int64(len(books)), Accessed: time.Now()})
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	var book *models.Book

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")
	param := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		panic(err)
	}

	defer conn.Disconnect(context.Background())

	filter := bson.D{{Key: "_id", Value: _id}}
	if err := coll.FindOne(context.Background(), filter).Decode(&book); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(SingleBook{Data: book, Length: 1, Accessed: time.Now()})
}
