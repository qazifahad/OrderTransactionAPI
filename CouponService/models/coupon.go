package models

import (
	// "reflect"
	// "errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Coupon struct {
    Id              bson.ObjectId   `json:"id" bson:"_id"`
    Name            string          `json:"name" bson:"name" validate:"nonzero"`
    ValidStartDate  time.Time       `json:"validStartDate" bson:"validStartDate" validate:"nonzero"`
    ValidEndDate    time.Time       `json:"validEndDate" bson:"validEndDate" validate:"nonzero"`
    Quantity        int             `json:"quantity" bson:"quantity"`
    DiscType        DiscountType    `json:"discType" bson:"discType" validate:"min=1,max=2"`
    DiscVal         int             `json:"discVal" bson:"discVal" validate:"nonzero"`
}

// func TimeValid(start,end interface{}, param time.Time) error {
//     startTime := reflect.ValueOf(start)
//     endTime := reflect.ValueOf(end)
// 	if startTime.Kind() != reflect.time{
// 		return errors.New("notZZ only validates strings")
// 	}
// 	if startTime.String() == "ZZ" {
// 		return errors.New("value cannot be ZZ")
// 	}
// 	return nil
// }