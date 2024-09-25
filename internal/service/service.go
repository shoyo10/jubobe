package service

import (
	"context"
	"jubobe/internal/model"
	"jubobe/internal/repository"
)

type service struct {
	repo repository.Repositorier
}

// New creates a new instance of the service
func New(repo repository.Repositorier) Servicer {
	return &service{
		repo: repo,
	}
}

func (s *service) ListPatients(ctx context.Context, opt *model.PatientOption) ([]model.Patient, error) {
	return s.repo.ListPatients(ctx, opt)
}

func (s *service) CreateOrder(ctx context.Context, order *model.Order) error {
	return s.repo.CreateOrder(ctx, order)
}

func (s *service) UpdateOrder(ctx context.Context, opt *model.OrderOption, in model.UpdateOrderInput) error {
	return s.repo.UpdateOrder(ctx, opt, in)
}
