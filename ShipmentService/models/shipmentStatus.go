package models

type ShipmentStatus int 

const (
	Manifested ShipmentStatus = iota
	OnTransit
	OnProcess
	ReceivedOnDestination
	Delivered
)