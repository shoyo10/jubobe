package pg

import (
	"context"
	"jubobe/internal/model"
	"jubobe/pkg/errors"
)

func (r *repo) ListPatients(ctx context.Context, opt *model.PatientOption) ([]model.Patient, error) {
	var patients []model.Patient
	db := r.Ctx(ctx)
	if opt != nil {
		db = db.Scopes(opt.Preload)
	}
	err := db.Find(&patients).Error
	if err != nil {
		return nil, errors.Wrapf(errors.ErrInternalServerError, "%v", err)
	}
	return patients, err
}
