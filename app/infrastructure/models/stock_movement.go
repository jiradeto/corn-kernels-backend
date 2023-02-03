package models

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
)

// StockMovement models
type StockMovement struct {
	ID           *string
	ProductID    *string
	Description  *string
	MovementType *string
	Quantity     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// TableName is table name of StockMovement
func (StockMovement) TableName() string {
	return "stock_movements"
}

// FromEntity converts entity to model
func (w *StockMovement) FromEntity(e *entities.StockMovement) (*StockMovement, error) {
	var model StockMovement

	err := copier.Copy(&model, &e)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// ToEntity converts model to entity
func (w *StockMovement) ToEntity() (*entities.StockMovement, error) {
	var entity entities.StockMovement

	err := copier.Copy(&entity, &w)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// StockMovements is an array of StockMovement
type StockMovements []*StockMovement

// FromEntities converts models to entities
func (ws StockMovements) FromEntities(es []*entities.StockMovement) ([]*StockMovement, error) {
	var ms []*StockMovement

	for _, e := range es {
		m, err := new(StockMovement).FromEntity(e)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

// ToEntities converts entities to models
func (ws StockMovements) ToEntities() (entities.StockMovements, error) {
	var es entities.StockMovements

	for _, w := range ws {
		e, err := w.ToEntity()
		if err != nil {
			return nil, err
		}
		es = append(es, e)
	}

	return es, nil
}
