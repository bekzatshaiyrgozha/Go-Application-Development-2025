package user

import (
	"context"
	"log"

	"w10/pkg/reqresp"
)

func WithLogging(service Service) Service {
	return &middlewareLogging{
		next: service,
	}
}

type middlewareLogging struct {
	next Service
}

func (s *middlewareLogging) GetUserByID(
	ctx context.Context,
	req reqresp.GetUserByIDRequest,
) (resp reqresp.GetUserByIDResponse, err error) {
	log.Printf("request: %v", req)
	defer func() {
		if err != nil {
			log.Printf("error: %v", err)
			return
		}

		log.Printf("response: %v", resp)
	}()

	return s.next.GetUserByID(ctx, req)
}
