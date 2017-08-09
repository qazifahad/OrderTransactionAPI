package models

type Order struct {
    Id              string          `json:"id"`
    CustomerId      string          `json:"customerId"`
    CouponId        string          `json:"couponId"`
    DeliveryInfo    Delivery        `json:"deliveryInfo"`
    PaymentInfo     Payment         `json:"paymentInfo"`
    Status          OrderStatus     `json:"status"`
}
