package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	var Books []*models.Book

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")

	defer conn.Disconnect(context.Background())
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.Background(), &Books); err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(BookList{Data: Books, Length: int64(len(Books)), Accessed: time.Now()})
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	var Book models.Book

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")
	param := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		panic(err)
	}

	defer conn.Disconnect(context.Background())

	filter := bson.D{{Key: "_id", Value: _id}}
	if err := coll.FindOne(context.Background(), filter).Decode(&Book); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(SingleBook{Data: &Book, Length: 1, Accessed: time.Now()})
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")
	defer conn.Disconnect(context.Background())

	if err := validateBookRequest(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.Id = primitive.NewObjectID()
	_, err := coll.InsertOne(context.Background(), book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error %v", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]string{"message": "object created"})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	return
}

func validateBookRequest(book *models.Book) error {
	if book.Title == "" || book.Author == "" || book.Page == 0 {
		return errors.New("missing required fields")
	}
	return nil
}
