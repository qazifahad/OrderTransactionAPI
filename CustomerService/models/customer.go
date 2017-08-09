package models

import (
	// Third party Libraries
	"gopkg.in/mgo.v2/bson"
)

type Customer struct {
    Id      bson.ObjectId   `json:"id" bson:"_id"`
    Name    string          `json:"name" bson:"name"`
}