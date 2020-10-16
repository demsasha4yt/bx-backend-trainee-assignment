package service

import (
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/emailsender"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"
)

// service ...
type service struct {
	store               store.Store
	emailSender         emailsender.EmailSender
	avitoAPI            avitoapi.AvitoAPI
	subscriptionService *subscriptionService
	emailsService       *emailsService
	adsService          *adsService
	priceTrackerService *priceTrackerService
	adsEmailsService    *adsEmailsService
}

// New ...
func New(store store.Store, emailsender emailsender.EmailSender,
	avitoAPI avitoapi.AvitoAPI) Service {
	return &service{
		store:       store,
		avitoAPI:    avitoAPI,
		emailSender: emailsender,
	}
}

// Subscription ...
func (s *service) Subscription() SubscriptionService {
	if s.subscriptionService != nil {
		return s.subscriptionService
	}
	s.subscriptionService = &subscriptionService{
		service:     s,
		avitoAPI:    s.avitoAPI,
		emailSender: s.emailSender,
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

func (s *service) PriceTracker() PriceTrackerService {
	if s.priceTrackerService != nil {
		return s.priceTrackerService
	}
	s.priceTrackerService = &priceTrackerService{
		service:     s,
		avitoAPI:    s.avitoAPI,
		emailSender: s.emailSender,
	}
	return s.priceTrackerService
}

func (s *service) AdsEmails() AdsEmailsService {
	if s.adsEmailsService != nil {
		return s.adsEmailsService
	}
	s.adsEmailsService = &adsEmailsService{
		service: s,
	}
	return s.adsEmailsService
}
