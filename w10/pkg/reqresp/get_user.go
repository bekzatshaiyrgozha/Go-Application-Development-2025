package reqresp

import "w10/pkg/domain"

type GetUserByIDRequest struct {
	ID string `json:"id"`
}

type GetUserByIDResponse struct {
	User domain.User `json:"user"`
}
