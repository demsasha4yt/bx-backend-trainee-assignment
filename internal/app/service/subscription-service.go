package service

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// SubscriptionService service
type SubscriptionService struct {
	service *Service
}

// Subscribe to ad
// Returns message and error
func (s *SubscriptionService) Subscribe(ctx context.Context, email string, link string) (interface{}, error) {
	ad := &models.Ad{
		Emails: []*models.Email{
			&models.Email{
				Email:     email,
				Confirmed: false,
			},
		},
	}
	return ad, nil
}
