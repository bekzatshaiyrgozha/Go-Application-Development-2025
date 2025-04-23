package pg

import (
	"context"

	"w10/internal/repositories/user"
	"w10/pkg/domain"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) user.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(ctx context.Context, id string) (domain.User, error) {
	return domain.User{
		ID:   id,
		Name: "Name" + id,
	}, nil
}
