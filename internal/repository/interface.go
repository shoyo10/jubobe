package repository

import (
	"context"
	"jubobe/internal/model"

	"gorm.io/gorm"
)

type Repositorier interface {
	PatientRepo
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
