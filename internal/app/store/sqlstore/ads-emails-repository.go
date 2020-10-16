package sqlstore

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"
	"github.com/jackc/pgx/v4"
)

// AdsEmailsRepository repository
type AdsEmailsRepository struct {
	store *Store
}

// FindByIds find model by ad and email ids
func (r *AdsEmailsRepository) FindByIds(ctx context.Context, adID, emailID int64) (*models.AdsEmails, error) {
	u := &models.AdsEmails{}

	if err := r.store.db.QueryRow(
		ctx,
		`SELECT id, ad_id, email_id, confirm, confirm_token, unsubscribe_token FROM ads_emails
		WHERE ad_id=$1 AND email_id=$2
		LIMIT 1`,
		adID, emailID,
	).Scan(
		&u.ID,
		&u.AdID,
		&u.EmailID,
		&u.Confirmed,
		&u.ConfirmToken,
		&u.UnsubscribeToken,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

// UpdateConfirmed update confirmed
func (r *AdsEmailsRepository) UpdateConfirmed(ctx context.Context, u *models.AdsEmails, confirmed bool) error {
	_, err := r.store.db.Exec(
		ctx,
		`UPDATE ads_emails SET confirm = $1 WHERE id=$2`,
		confirmed,
		u.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes relationship
func (r *AdsEmailsRepository) Delete(ctx context.Context, u *models.AdsEmails) error {
	_, err := r.store.db.Exec(
		ctx,
		`DELETE FROM ads_emails WHERE id=$1`,
		u.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
