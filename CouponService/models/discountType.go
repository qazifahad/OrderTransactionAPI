package models

type DiscountType int

const (
	Percentage DiscountType = iota
	Nominal
)