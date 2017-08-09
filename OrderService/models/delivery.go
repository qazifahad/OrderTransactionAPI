package models

type Delivery struct {
	Name 		string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email 		string `json:"email"`
	Address 	string `json:"address"`
}