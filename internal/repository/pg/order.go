package pg

import (
	"context"
	"jubobe/internal/model"
	"jubobe/pkg/errors"

	"gorm.io/gorm"
)

func (r *repo) CreateOrder(ctx context.Context, order *model.Order) error {
	err := r.Ctx(ctx).Omit("id").Create(order).Error
	if err != nil {
		return errors.Wrapf(errors.ErrInternalServerError, "%v", err)
	}
	return nil
}

func (r *repo) UpdateOrder(ctx context.Context, opt *model.OrderOption, in model.UpdateOrderInput) error {
	db := r.Ctx(ctx)
	if opt != nil {
		db = opt.Filter.Where(db)
	}
	result := db.Model(&model.Order{}).Updates(in)
	rawEffected, err := result.RowsAffected, result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrapf(errors.ErrResourceNotFound, "%v", err)
		}
		return errors.Wrapf(errors.ErrInternalServerError, "%v", err)
	}
	if rawEffected == 0 {
		return errors.WithStack(errors.ErrResourceNotFound)
	}
	return nil
}
