package pg

import (
	"context"
	"jubobe/internal/model"
	"jubobe/pkg/errors"
)

func (r *repo) CreateOrder(ctx context.Context, order *model.Order) error {
	err := r.Ctx(ctx).Omit("id").Create(order).Error
	return errors.Wrapf(errors.ConvertPostgresError(err), "%v", err)
}

func (r *repo) UpdateOrder(ctx context.Context, opt *model.OrderOption, in model.UpdateOrderInput) error {
	db := r.Ctx(ctx)
	if opt != nil {
		db = opt.Filter.Where(db)
	}
	result := db.Model(&model.Order{}).Updates(in)
	rawEffected, err := result.RowsAffected, result.Error
	if err != nil {
		errors.Wrapf(errors.ConvertPostgresError(err), "%v", err)
	}
	if rawEffected == 0 {
		return errors.WithStack(errors.ErrResourceNotFound)
	}
	return nil
}
