package store

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// AdsRepository interface
type AdsRepository interface {
	Create(context.Context, *models.Ad) error
	FindAll(context.Context, int, int) ([]*models.Ad, error)
}

// EmailsRepository interface
type EmailsRepository interface {
	Create(context.Context, *models.Email) error
	FindByEmail(context.Context, string) (*models.Email, error)
}
