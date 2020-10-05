package service

import "github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"

// Service ...
type Service struct {
	store               store.Store
	subscriptionService *SubscriptionService
}

// New ...
func New(store store.Store) Interface {
	return &Service{
		store: store,
	}
}

// Subscription ...
func (s *Service) Subscription() SubscriptionServiceInterface {
	if s.subscriptionService != nil {
		return s.subscriptionService
	}
	s.subscriptionService = &SubscriptionService{
		service: s,
	}
	return s.subscriptionService
}
