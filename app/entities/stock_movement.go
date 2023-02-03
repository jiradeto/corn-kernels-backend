package entities

import "time"

// StockMovement entity
type StockMovement struct {
	CreatedAt    *time.Time
	ID           *string
	ProductID    *string
	Description  *string
	MovementType *string
	Quantity     int
}

// StockMovements is an array of type StockMovement
type StockMovements []*StockMovement
