package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/PetrusAriaa/Go-HTTP-Server/lib"
	"github.com/PetrusAriaa/Go-HTTP-Server/logger"
	"github.com/PetrusAriaa/Go-HTTP-Server/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookList struct {
	Data     []models.Book `json:"data"`
	Length   int64         `json:"length"`
	Accessed time.Time     `json:"accessed"`
}

type SingleBook struct {
	Data     models.Book `json:"data"`
	Length   int64       `json:"length"`
	Accessed time.Time   `json:"accessed"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	var Books []models.Book

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")
	defer conn.Disconnect(context.Background())

	c, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		logger.ResponseLogger(r, http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = c.All(context.Background(), &Books); err != nil {
		logger.ResponseLogger(r, http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(BookList{Data: Books, Length: int64(len(Books)), Accessed: time.Now()})
	logger.ResponseLogger(r, http.StatusOK, "Response sent")
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	var Book models.Book

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")
	defer conn.Disconnect(context.Background())

	param := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		logger.ResponseLogger(r, http.StatusBadRequest, err.Error())
		http.Error(w, "server: bad request", http.StatusBadRequest)
		return
	}

	filter := bson.D{{Key: "_id", Value: _id}}
	if err := coll.FindOne(context.Background(), filter).Decode(&Book); err != nil {
		if err == mongo.ErrNoDocuments {
			logger.ResponseLogger(r, http.StatusNotFound, err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	// json.NewEncoder(w).Encode(SingleBook{Data: Book, Length: 1, Accessed: time.Now()})
	json.NewEncoder(w).Encode(map[string]string{"message": "Test GHAction"})
	logger.ResponseLogger(r, http.StatusOK, "Response sent")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var Book models.Book

	if err := json.NewDecoder(r.Body).Decode(&Book); err != nil {
		logger.ResponseLogger(r, http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")
	defer conn.Disconnect(context.Background())

	if err := validateBookRequest(&Book); err != nil {
		logger.ResponseLogger(r, http.StatusBadRequest, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Book.Id = primitive.NewObjectID()
	_, err := coll.InsertOne(context.Background(), Book)
	if err != nil {
		logger.ResponseLogger(r, http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]string{"message": "object created"})
	logger.ResponseLogger(r, http.StatusCreated, "Document created successfully")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	conn := lib.ConnectDB()
	coll := conn.Database("bookstore").Collection("book")
	defer conn.Disconnect(context.Background())

	param := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		logger.ResponseLogger(r, http.StatusBadRequest, err.Error())
		http.Error(w, "server: bad request", http.StatusBadRequest)
		return
	}

	res, err := coll.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: _id}})
	if err != nil {
		logger.ResponseLogger(r, http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.DeletedCount == 0 {
		logger.ResponseLogger(r, http.StatusNotFound, "mongo: no documents to delete")
		http.Error(w, "mongo: no documents to delete", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	logger.ResponseLogger(r, http.StatusNoContent, "Document deleted successfully")
}

func validateBookRequest(book *models.Book) error {
	if book.Metadata.Author == "" || book.Metadata.ISBN == "" || book.Metadata.Title == "" || book.Metadata.Page == 0 {
		return errors.New("server: missing required fields")
	}
	return nil
}
