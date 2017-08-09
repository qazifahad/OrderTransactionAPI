package models

type Payment struct {
	Code			string	`json:"code"`
	Name			string	`json:"name"`
	TransferAmount	int		`json:"transferAmount"`
}