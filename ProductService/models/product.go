package models

type Product struct {
    Id          string  `json:"id"`
    Name        string  `json:"name"`
    Quantity    int     `json:"quantity"`
    Price       int     `json:"price"`
}