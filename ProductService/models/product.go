package models

import (
    "gopkg.in/mgo.v2/bson"
)

type Product struct {
    Id          bson.ObjectId   `json:"id" bson:"_id"`                              // required, no defaults
    Name        string          `json:"name" bson:"name" validate:"nonzero"`        // required, no defaults
    Quantity    int             `json:"quantity" bson:"quantity"`                   // required, uses defaults : 0
    Price       int             `json:"price" bson:"price" validate:"nonzero"`      // required, no defaults
}