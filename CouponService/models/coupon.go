package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Coupon struct {
    Id              bson.ObjectId   `json:"id" bson:"_id"`
    Name            string          `json:"name" bson:"name"`
    ValidStartDate  time.Time       `json:"validStartDate" bson:"validStartDate"`
    ValidEndDate    time.Time       `json:"validEndDate" bson:"validEndDate"`
    Quantity        int             `json:"quantity" bson:"quantity"`
    DiscType        DiscountType    `json:"discType" bson:"discType"`
    DiscVal         int             `json:"discVal" bson:"discVal"`
}