package pg

import (
	"context"
	"jubobe/internal/repository"

	"gorm.io/gorm"
)

type repo struct {
	conn *gorm.DB
}

// New creates a new instance of the repository
func New(conn *gorm.DB) (repository.Repositorier, error) {
	r := &repo{
		conn: conn,
	}
	return r, nil
}

func (r *repo) Ctx(ctx context.Context) *gorm.DB {
	return r.conn.WithContext(ctx)
}
