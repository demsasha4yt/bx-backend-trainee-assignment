package sqlstore

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// AdsRepository repository
type AdsRepository struct {
	store *Store
}

// Create creates new ad in db
func (r *AdsRepository) Create(ctx context.Context, u *models.Ad) error {
	tx, err := r.store.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := tx.QueryRow(
		ctx,
		"INSERT INTO ads (avito_id, actual) VALUES ($1, $2) RETURNING id",
		u.AvitoID,
		u.Actual,
	).Scan(
		&u.ID,
	); err != nil {
		return err
	}

	if u.Emails != nil && len(u.Emails) > 0 {
		for _, email := range u.Emails {
			_, err := tx.Exec(
				ctx,
				"INSERT INTO ads_emails(ad_id, email_id) VALUES ($1, $2)",
				u.ID, email.ID,
			)
			if err != nil {
				return err
			}
		}
	}

	if _, err := tx.Exec(
		ctx,
		"INSERT INTO ads_prices (ad_id, price) VALUES($1, $2)",
		u.ID, u.CurrentPrice,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// FindAll finds all actual ads with confirmed email
func (r *AdsRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.Ad, error) {
	return nil, nil
}
