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
}

// SubscriptionService interface
type SubscriptionService interface {
	Subscribe(context.Context, string, string) (interface{}, error)
}

// AdsService interface
type AdsService interface {
	Create(context.Context, *models.Ad) error
}

// EmailsService interface
type EmailsService interface {
	Create(context.Context, *models.Email) error
	FindByEmail(context.Context, string) (*models.Email, error)
	FindByEmailOrCreate(context.Context, string) (*models.Email, error)
}
