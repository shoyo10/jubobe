package service

import (
	"context"
	"jubobe/internal/model"
)

type Servicer interface {
	ListPatients(ctx context.Context, opt *model.PatientOption) ([]model.Patient, error)
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, opt *model.OrderOption, in model.UpdateOrderInput) error
}
