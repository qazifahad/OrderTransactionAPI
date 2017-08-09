package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Shipment struct {
    Id          bson.ObjectId   `json:"id" bson:"_id"`
    OrderId     int             `json:"orderId" bson:"orderId"`
    Status      ShipmentStatus  `json:"status" bson:"status"`
}