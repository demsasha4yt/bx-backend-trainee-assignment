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

// SubscribeOrCreate subscribe Ad.Emails or Create new ad in db..
func (s *adsService) SubscribeOrCreate(ctx context.Context, u *models.Ad) error {
	ad, err := s.service.store.Ads().FindByAvitoID(ctx, u.AvitoID)
	if err != nil {
		if err := s.service.store.Ads().Create(ctx, u); err != nil {
			return err
		}
		return nil
	}
	u.ID = ad.ID
	if err := s.service.store.Ads().CreateEmails(ctx, u); err != nil {
		return err
	}
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
