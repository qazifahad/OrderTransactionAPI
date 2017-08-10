package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
    Id              bson.ObjectId   `json:"id" bson:"_id"`
    CustomerId      string          `json:"customerId" bson:"customerId"`
    CouponId        string          `json:"couponId" bson:"couponId"`
    DeliveryInfo    Delivery        `json:"deliveryInfo" bson:"deliveryInfo"`
    PaymentInfo     Payment         `json:"paymentInfo" bson:"paymentInfo"`
    Status          OrderStatus     `json:"status" bson:"status"`
    OrderLines      []OrderLine     `json:"orderLines" bson:"orderLines"`
}
