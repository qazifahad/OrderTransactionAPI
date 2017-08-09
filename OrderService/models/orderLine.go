package models

type OrderLine struct {
	OrderId 	string 	`json:"orderId"`
	ProductId 	string 	`json:"productId"`
	Quantity 	int 	`json:"quantity"`
}
