package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id       primitive.ObjectID `bson:"_id"`
	Metadata Metadata           `json:"metadata"`
	Rate     int                `json:"rate"`
	Price    int                `json:"price"`
	Stocks   int                `json:"stocks"`
}

type Metadata struct {
	ISBN   string `json:"ISBN"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Page   int    `json:"page"`
}
