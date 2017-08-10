package models

type DiscountType int

const (
	Percentage DiscountType = iota + 1
	Nominal
)