package entities

import "time"

// Product entity
type Product struct {
	CreatedAt   *time.Time
	ID          *string
	Name        *string
	Description *string
	Quantity    int
}

// Products is an array of type Product
type Products []*Product
