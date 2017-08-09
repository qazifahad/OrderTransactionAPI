package models

type Shipment struct {
    Id          int             `json:"id"`
    OrderId     int             `json:"orderId"`
    Status      ShipmentStatus  `json:"status"`
}