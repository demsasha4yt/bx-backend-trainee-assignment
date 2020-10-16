package service

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

type adsEmailsService struct {
	service *service
}

// FindByIds find model by ad and email ids
func (s *adsEmailsService) FindByIds(ctx context.Context, adID, emailID int64) (*models.AdsEmails, error) {
	return s.service.store.AdsEmails().FindByIds(ctx, adID, emailID)
}

// UpdateConfirmed update confirmed
func (s *adsEmailsService) UpdateConfirmed(ctx context.Context, u *models.AdsEmails, confirmed bool) error {
	return s.service.store.AdsEmails().UpdateConfirmed(ctx, u, confirmed)
}

// Delete deletes relationship
func (s *adsEmailsService) Delete(ctx context.Context, u *models.AdsEmails) error {
	return s.service.store.AdsEmails().Delete(ctx, u)
}
