package service

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// SubscriptionService service
type SubscriptionService struct {
	service *Service
}

// Subscribe to ad
// Returns message and error
func (s *SubscriptionService) Subscribe(ctx context.Context, email string, link string) (string, error) {
	if err := validation.Validate(&email, validation.Required, is.Email); err != nil {
		return "", err
	}
	ad := &models.Ad{}
	if err := s.service.store.Ads().Create(ctx, ad); err != nil {
		return "", err
	}
	if err := ad.SetAvitoIDFromLink(link); err != nil {
		return "", err
	}
	return "OK", nil
}
