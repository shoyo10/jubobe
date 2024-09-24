package repository

import (
	"context"
	"jubobe/internal/model"

	"gorm.io/gorm"
)

type Repositorier interface {
	PatientRepo
	OrderRepo
}

type PatientRepo interface {
	ListPatients(ctx context.Context, opt *PatientOption) ([]model.Patient, error)
}

type PatientOption struct {
	IsPreloadOrder bool
}

func (o *PatientOption) Preload(db *gorm.DB) *gorm.DB {
	if o.IsPreloadOrder {
		db = db.Preload("Order")
	}
	return db
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, opt *OrderOption, in UpdateOrderInput) error
}

type OrderOption struct {
	Filter OrderFilter
}

type OrderFilter struct {
	ID int
}

func (o *OrderFilter) Where(db *gorm.DB) *gorm.DB {
	if o.ID != 0 {
		db = db.Where("id = ?", o.ID)
	}
	return db
}

type UpdateOrderInput struct {
	Message string
}
