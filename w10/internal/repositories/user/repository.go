package user

import (
	"context"

	"w10/pkg/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (domain.User, error)
}
