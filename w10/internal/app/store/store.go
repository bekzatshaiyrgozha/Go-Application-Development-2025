package store

import (
	"w10/internal/app/connections"
	"w10/internal/repositories/user"
	"w10/internal/repositories/user/pg"
)

type Store struct {
	UserRepository user.Repository
}

func NewStore(conns *connections.Connections) *Store {
	st := &Store{
		UserRepository: pg.New(conns.DB),
	}

	return st
}
