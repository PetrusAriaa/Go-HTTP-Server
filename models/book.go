package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id     primitive.ObjectID `bson:"_id"`
	Title  string
	Author string
	Page   int
	Rate   int
}
