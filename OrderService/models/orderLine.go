package models

import (
	"gopkg.in/mgo.v2/bson"
)

type OrderLine struct {
	ProductId 	bson.ObjectId 	`json:"productId" bson:"productId"`
	Quantity 	int 			`json:"quantity" bson:"quantity"`
}
