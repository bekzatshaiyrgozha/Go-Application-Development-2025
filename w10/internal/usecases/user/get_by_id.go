package user

import (
	"context"

	"w10/internal/app/store"
	"w10/pkg/domain"
	"w10/pkg/reqresp"
)

type GetUserByIDRepository interface {
	GetUserByID(ctx context.Context, id string) (domain.User, error)
}

func GetUserByID(ctx context.Context, repo GetUserByIDRepository, req reqresp.GetUserByIDRequest) (reqresp.GetUserByIDResponse, error) {
	user, err := repo.GetUserByID(ctx, req.ID)
	if err != nil {
		return reqresp.GetUserByIDResponse{}, err
	}

	return reqresp.GetUserByIDResponse{User: user}, nil
}

func NewGetUserByIDRepository(st *store.Store) GetUserByIDRepository {
	return &getUserByIDRepositoryFacade{st: st}
}

type getUserByIDRepositoryFacade struct {
	st *store.Store
}

func (f *getUserByIDRepositoryFacade) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	return f.st.UserRepository.GetByID(ctx, id)
}
