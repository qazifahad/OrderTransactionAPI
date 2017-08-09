package models

import (
	"time"
)

type Coupon struct {
    Id              string      `json:"Id"`
    Name            string      `json:"name"`
    ValidStartDate  time.Time   `json:"validStartDate"`
    ValidEndDate    time.Time   `json:"validEndDate"`
    Quantity        int         `json:"quantity"`
    DiscType        string      `json:"discType"`
    DiscVal         int         `json:"discVal"`
}