package service

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// Service interface
type Service interface {
	Subscription() SubscriptionService
	Emails() EmailsService
	Ads() AdsService
	AdsEmails() AdsEmailsService
	PriceTracker() PriceTrackerService
}

// SubscriptionService interface
type SubscriptionService interface {
	Subscribe(context.Context, string, string) (interface{}, error)
	ConfirmSubscribe(context.Context, int64, int64, string) (interface{}, error)
	Unsubscribe(context.Context, int64, int64, string) (interface{}, error)
}

// AdsService interface
type AdsService interface {
	Create(context.Context, *models.Ad) error
	SubscribeOrCreate(context.Context, *models.Ad) error
}

// EmailsService interface
type EmailsService interface {
	Create(context.Context, *models.Email) error
	FindByEmail(context.Context, string) (*models.Email, error)
	FindByEmailOrCreate(context.Context, string) (*models.Email, error)
}

// PriceTrackerService interface
type PriceTrackerService interface {
	CheckAdsTask() error
}

// AdsEmailsService interface
type AdsEmailsService interface {
	FindByIds(context.Context, int64, int64) (*models.AdsEmails, error)
	UpdateConfirmed(context.Context, *models.AdsEmails, bool) error
	Delete(context.Context, *models.AdsEmails) error
}
