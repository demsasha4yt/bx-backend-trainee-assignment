package store

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// AdsRepository interface
type AdsRepository interface {
	Create(context.Context, *models.Ad) error
	FindByAvitoID(context.Context, int64) (*models.Ad, error)
	CreateEmails(context.Context, *models.Ad) error
	FindAll(context.Context, int, int) ([]*models.Ad, error)
}

// EmailsRepository interface
type EmailsRepository interface {
	Create(context.Context, *models.Email) error
	FindByEmail(context.Context, string) (*models.Email, error)
}

// AdsEmailsRepository interface
type AdsEmailsRepository interface {
	FindByIds(context.Context, int64, int64) (*models.AdsEmails, error)
	UpdateConfirmed(context.Context, *models.AdsEmails, bool) error
	Delete(context.Context, *models.AdsEmails) error
}
