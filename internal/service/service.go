package service

import (
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
