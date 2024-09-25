package errors

import (
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func ConvertPostgresError(err error) error {
	if err == nil {
		return nil
	}
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case "23505":
			return ErrResourceAlreadyExists
		}
	}

	if Is(err, gorm.ErrRecordNotFound) {
		return ErrResourceNotFound
	}
	return ErrInternalServerError
}
