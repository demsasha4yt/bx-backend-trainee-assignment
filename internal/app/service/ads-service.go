package service

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

type adsService struct {
	service *service
}

// Create creates new ad
func (s *adsService) Create(ctx context.Context, u *models.Ad) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := s.service.store.Ads().Create(ctx, u); err != nil {
		return err
	}
	// Something with Ad ...
	return nil
}

// Create finds actual ads
func (s *adsService) FindAll(ctx context.Context, offset, limit int) ([]*models.Ad, error) {
	ads, err := s.service.store.Ads().FindAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return ads, nil
}
