package service

import "github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"

// service ...
type service struct {
	store               store.Store
	subscriptionService *subscriptionService
	emailsService       *emailsService
	adsService          *adsService
}

// New ...
func New(store store.Store) Service {
	return &service{
		store: store,
	}
}

// Subscription ...
func (s *service) Subscription() SubscriptionService {
	if s.subscriptionService != nil {
		return s.subscriptionService
	}
	s.subscriptionService = &subscriptionService{
		service: s,
	}
	return s.subscriptionService
}

func (s *service) Emails() EmailsService {
	if s.emailsService != nil {
		return s.emailsService
	}
	s.emailsService = &emailsService{
		service: s,
	}
	return s.emailsService
}

func (s *service) Ads() AdsService {
	if s.adsService != nil {
		return s.adsService
	}
	s.adsService = &adsService{
		service: s,
	}
	return s.adsService
}
