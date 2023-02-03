package models

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
)

// ProductStock models
type ProductStock struct {
	ID        *string
	ProductID *string
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName is table name of ProductStock
func (ProductStock) TableName() string {
	return "product_stocks"
}

// FromEntity converts entity to model
func (w *ProductStock) FromEntity(e *entities.ProductStock) (*ProductStock, error) {
	var model ProductStock

	err := copier.Copy(&model, &e)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// ToEntity converts model to entity
func (w *ProductStock) ToEntity() (*entities.ProductStock, error) {
	var entity entities.ProductStock

	err := copier.Copy(&entity, &w)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// ProductStocks is an array of ProductStock
type ProductStocks []*ProductStock

// FromEntities converts models to entities
func (ws ProductStocks) FromEntities(es []*entities.ProductStock) ([]*ProductStock, error) {
	var ms []*ProductStock

	for _, e := range es {
		m, err := new(ProductStock).FromEntity(e)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

// ToEntities converts entities to models
func (ws ProductStocks) ToEntities() (entities.ProductStocks, error) {
	var es entities.ProductStocks

	for _, w := range ws {
		e, err := w.ToEntity()
		if err != nil {
			return nil, err
		}
		es = append(es, e)
	}

	return es, nil
}
