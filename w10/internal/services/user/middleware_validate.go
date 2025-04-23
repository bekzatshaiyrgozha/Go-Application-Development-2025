package user

import (
	"context"

	"w10/pkg/reqresp"
)

func WithValidate(service Service) Service {
	return &middlewareValidate{
		next: service,
	}
}

type middlewareValidate struct {
	next Service
}

func (s *middlewareValidate) GetUserByID(
	ctx context.Context,
	req reqresp.GetUserByIDRequest,
) (resp reqresp.GetUserByIDResponse, err error) {
	// todo: validation

	return s.next.GetUserByID(ctx, req)
}
