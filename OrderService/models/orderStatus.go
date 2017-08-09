package models

type OrderStatus int

const (
	AwaitingPayment OrderStatus = iota
	AwaitingShipment
	Shipped
	Canceled
)