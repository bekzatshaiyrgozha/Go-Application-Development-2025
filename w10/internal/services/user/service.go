package user

import (
	"context"

	"w10/internal/app/store"
	"w10/internal/usecases/user"
	"w10/pkg/middleware"
	"w10/pkg/reqresp"
)

type Service interface {
	GetUserByID(ctx context.Context, req reqresp.GetUserByIDRequest) (reqresp.GetUserByIDResponse, error)
}

type service struct {
	st *store.Store
}

func NewService(st *store.Store) (srv Service) {
	srv = &service{st: st}
	srv = WithValidate(srv)
	srv = WithLogging(srv)

	return srv
}

func (s service) GetUserByID(
	ctx context.Context,
	req reqresp.GetUserByIDRequest,
) (resp reqresp.GetUserByIDResponse, err error) {
	err = middleware.WaitContextCancel(ctx, "GetUserByID", func() error {
		var useCaseErr error
		resp, useCaseErr = user.GetUserByID(ctx, user.NewGetUserByIDRepository(s.st), req)

		return useCaseErr
	})

	return
}
