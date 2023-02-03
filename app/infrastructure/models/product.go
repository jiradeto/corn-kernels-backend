package models

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
)

// Product models
type Product struct {
	ID           string
	Name         string
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProductStock *ProductStock
}

// TableName is table name of Product
func (Product) TableName() string {
	return "products"
}

// FromEntity converts entity to model
func (w *Product) FromEntity(e *entities.Product) (*Product, error) {
	var model Product

	err := copier.Copy(&model, &e)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// ToEntity converts model to entity
func (w *Product) ToEntity() (*entities.Product, error) {
	var entity entities.Product

	err := copier.Copy(&entity, &w)
	if err != nil {
		return nil, err
	}

	if w.ProductStock != nil {
		entity.Quantity = w.ProductStock.Quantity
	}

	return &entity, nil
}

// Products is an array of Product
type Products []*Product

// FromEntities converts models to entities
func (ws Products) FromEntities(es []*entities.Product) ([]*Product, error) {
	var ms []*Product

	for _, e := range es {
		m, err := new(Product).FromEntity(e)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

// ToEntities converts entities to models
func (ws Products) ToEntities() (entities.Products, error) {
	var es entities.Products

	for _, w := range ws {
		e, err := w.ToEntity()
		if err != nil {
			return nil, err
		}
		es = append(es, e)
	}

	return es, nil
}
