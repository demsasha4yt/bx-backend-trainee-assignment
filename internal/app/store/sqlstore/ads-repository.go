package sqlstore

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"
	"github.com/jackc/pgx/v4"
)

// AdsRepository repository
type AdsRepository struct {
	store *Store
}

func txAppendEmails(ctx context.Context, tx pgx.Tx, u *models.Ad) error {
	if u.Emails != nil && len(u.Emails) > 0 {
		for _, email := range u.Emails {
			_, err := tx.Exec(
				ctx,
				`INSERT INTO ads_emails(ad_id, email_id, confirm_token, unsubscribe_token) 
					VALUES ($1, $2, $3, $4)`,
				u.ID, email.ID, email.ConfirmToken, email.UnsubscribeToken,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

	if err := txAppendEmails(ctx, tx, u); err != nil {
		return err
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

// CreateEmails append emails to ad in ads_emails table
func (r *AdsRepository) CreateEmails(ctx context.Context, u *models.Ad) error {
	tx, err := r.store.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err := txAppendEmails(ctx, tx, u); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// FindByAvitoID find ad by AvitoID
func (r *AdsRepository) FindByAvitoID(ctx context.Context, avitoID int64) (*models.Ad, error) {
	u := &models.Ad{}
	err := r.store.db.QueryRow(
		ctx,
		`SELECT id, actual, current_price FROM ads WHERE avito_id = $1`,
		avitoID,
	).Scan(
		&u.ID,
		&u.Actual,
		&u.CurrentPrice,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

// FindAll finds all actual ads with confirmed email
func (r *AdsRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.Ad, error) {
	ads := make([]*models.Ad, 0)

	rows, err := r.store.db.Query(
		ctx,
		`SELECT a.id, a.avito_id, a.current_price, a.actual, 
			COALESCE(json_agg(e) FILTER (WHERE e.id IS NOT NULL), '[]') AS emails
		FROM ads a
		LEFT JOIN ads_emails ae ON ae.ad_id = a.id
		LEFT JOIN emails e ON e.id = ae.email_id
		WHERE a.actual = true
		GROUP BY a.id
		OFFSET $1 LIMIT $2`,
		offset, limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := &models.Ad{}
		var emails []byte

		if err := rows.Scan(
			&u.ID,
			&u.AvitoID,
			&u.CurrentPrice,
			&u.Actual,
			&emails,
		); err != nil {
			return nil, err
		}

		emailsSlice, err := models.NewEmailSliceFromByte(emails)
		if err == nil && emailsSlice != nil {
			u.Emails = emailsSlice
		}

		ads = append(ads, u)
	}

	return ads, nil
}
