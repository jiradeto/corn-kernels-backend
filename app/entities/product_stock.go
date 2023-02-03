package entities

import "time"

// ProductStock entity
type ProductStock struct {
	CreatedAt *time.Time
	ID        *string
	ProductID *string
	Quantity  int
}

// ProductStocks is an array of type ProductStock
type ProductStocks []*ProductStock
