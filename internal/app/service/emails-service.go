package service

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"
)

type emailsService struct {
	service *service
}

// Create creates new email
func (s *emailsService) Create(ctx context.Context, u *models.Email) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := s.service.store.Emails().Create(ctx, u); err != nil {
		return err
	}
	return nil
}

// FindByEmail finds email by its value
func (s *emailsService) FindByEmail(ctx context.Context, email string) (*models.Email, error) {
	e, err := s.service.store.Emails().FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// FindOrCreate find email or creates
func (s *emailsService) FindByEmailOrCreate(ctx context.Context, email string) (*models.Email, error) {
	e, err := s.FindByEmail(ctx, email)

	if err != nil {
		if err == store.ErrRecordNotFound {
			e = &models.Email{
				Email: email,
			}
			if err := s.Create(ctx, e); err != nil {
				return nil, err
			}
			return e, nil
		}
		return nil, err
	}
	return e, nil
}
