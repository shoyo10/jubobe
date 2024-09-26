package repository

import (
	"context"
	"jubobe/internal/model"
)

type Repositorier interface {
	PatientRepo
	OrderRepo
}

type PatientRepo interface {
	ListPatients(ctx context.Context, opt *model.PatientOption) ([]model.Patient, error)
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, opt *model.OrderOption, in model.UpdateOrderInput) error
	GetOrder(ctx context.Context, opt *model.OrderOption) (*model.Order, error)
}
