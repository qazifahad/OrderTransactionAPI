package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
    Id          bson.ObjectId   `json:"id" bson:"_id"`
    Name        string          `json:"name" bson:"name"`
    Quantity    int             `json:"quantity" bson:"quantity"`
    Price       int             `json:"price" bson:"price"`
}