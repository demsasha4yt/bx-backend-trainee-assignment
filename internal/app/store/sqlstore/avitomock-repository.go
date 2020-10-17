package sqlstore

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// AvitoMockRepository repository
type AvitoMockRepository struct {
	store *Store
}

// Create creates new ad
func (r *AvitoMockRepository) Create(ctx context.Context, u *models.AvitoMock) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindByAvitoID returns ad in table or nil
func (r *AvitoMockRepository) FindByAvitoID(ctx context.Context, avitoID int64) (*models.AvitoMock, error) {
	return nil, nil
}

// UpdatePrices updates all prices to random value where count_changes < 3 && checked_at < time.Now - 10 min
func (r *AvitoMockRepository) UpdatePrices(ctx context.Context) error {
	return nil
}

// UpdatePrice for ad with AvitoID
func (r *AvitoMockRepository) UpdatePrice(ctx context.Context, avitoID int64, price int) error {
	return nil
}

// SetDeletedOne sets delete column to ad with avitoID
func (r *AvitoMockRepository) SetDeletedOne(ctx context.Context, avitoID int64) error {
	return nil
}

// SetDeleted sets delete column where count_changes > 2 && checked_at < time.Now - 10 min
func (r *AvitoMockRepository) SetDeleted(ctx context.Context) error {
	return nil
}
