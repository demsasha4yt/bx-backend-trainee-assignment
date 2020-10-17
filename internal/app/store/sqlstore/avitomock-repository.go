package sqlstore

import (
	"context"
	"time"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

var timestampSub int = -10

// AvitoMockRepository repository
type AvitoMockRepository struct {
	store *Store
}

// Create creates new ad
func (r *AvitoMockRepository) Create(ctx context.Context, u *models.AvitoMock) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := r.store.db.QueryRow(ctx,
		`INSERT INTO avito_mockapi (avito_id, price) VALUES ($1, $2) RETURNING id`,
		u.AvitoID, u.Price,
	).Scan(
		&u.ID,
	); err != nil {
		return err
	}
	return nil
}

// FindByAvitoID returns ad in table or nil
func (r *AvitoMockRepository) FindByAvitoID(ctx context.Context, avitoID int64) (*models.AvitoMock, error) {
	u := &models.AvitoMock{}

	if err := r.store.db.QueryRow(ctx,
		`SELECT id, avito_id, price, deleted FROM avito_mockapi WHERE avito_id = $1 LIMIT 1`,
		avitoID).Scan(
		&u.ID,
		&u.AvitoID,
		&u.Price,
		&u.Deleted,
	); err != nil {
		return nil, err
	}

	return u, nil
}

// UpdatePrices updates all prices to random value where count_changes < 3 && changed_at < time.Now - 10 min
func (r *AvitoMockRepository) UpdatePrices(ctx context.Context) error {
	timestamp := time.Now().Add(time.Duration(timestampSub) * time.Minute)

	_, err := r.store.db.Exec(ctx,
		`UPDATE avito_mockapi SET (price, count_changes) = (floor(random() * 10000 + 1), count_changes + 1)
		WHERE deleted != TRUE AND count_changes < 3 AND changed_at <= $1`,
		timestamp,
	)

	if err != nil {
		return err
	}

	return nil
}

// SetDeleted sets delete column where count_changes > 2 && changed_at < time.Now - 10 min
func (r *AvitoMockRepository) SetDeleted(ctx context.Context) error {
	timestamp := time.Now().Add(time.Duration(timestampSub) * time.Minute)

	_, err := r.store.db.Exec(ctx,
		`UPDATE avito_mockapi SET (price, count_changes) = (floor(random() * 10000 + 1), count_changes + 1)
		WHERE deleted != TRUE AND count_changes = 2 AND changed_at <= $1`,
		timestamp,
	)

	if err != nil {
		return err
	}

	return nil
}

// UpdatePrice for ad with AvitoID
func (r *AvitoMockRepository) UpdatePrice(ctx context.Context, avitoID int64, price int) error {
	_, err := r.store.db.Exec(ctx,
		`UPDATE avito_mockapi SET price = $1 WHERE avito_id = $2`,
		price, avitoID,
	)

	if err != nil {
		return err
	}

	return nil
}

// SetDeletedOne sets delete column to ad with avitoID
func (r *AvitoMockRepository) SetDeletedOne(ctx context.Context, avitoID int64) error {
	_, err := r.store.db.Exec(ctx,
		`UPDATE avito_mockapi SET deleted = TRUE WHERE avito_id = $1`,
		avitoID,
	)

	if err != nil {
		return err
	}

	return nil
}
