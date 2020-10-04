package service

import "github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"

// Service ...
type Service struct {
	store store.Store
}

// New ...
func New(store store.Store) Interface {
	return &Service{
		store: store,
	}
}
