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
	d := model.Order{
		Message: in.Message,
	}
	result := db.Model(&model.Order{}).Updates(d)
	rawEffected, err := result.RowsAffected, result.Error
	if err != nil {
		errors.Wrapf(errors.ConvertPostgresError(err), "%v", err)
	}
	if rawEffected == 0 {
		return errors.WithStack(errors.ErrResourceNotFound)
	}
	return nil
}

func (r *repo) GetOrder(ctx context.Context, opt *model.OrderOption) (*model.Order, error) {
	db := r.Ctx(ctx)
	if opt != nil {
		db = opt.Filter.Where(db)
	}
	var order model.Order
	err := db.First(&order).Error
	if err != nil {
		return nil, errors.Wrapf(errors.ConvertPostgresError(err), "%v", err)
	}
	return &order, nil
}
