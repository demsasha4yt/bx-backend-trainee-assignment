package sqlstore

import (
	"context"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// EmailsRepository repository
type EmailsRepository struct {
	store *Store
}

// Create creates new email in DB
func (r *EmailsRepository) Create(ctx context.Context, u *models.Email) error {
	return r.store.db.QueryRow(
		ctx,
		"INSERT INTO emails (email) VALUES ($1) RETURNING id",
		u.Email,
	).Scan(
		&u.ID,
	)
}

// FindByEmail finds email by its value
func (r *EmailsRepository) FindByEmail(ctx context.Context, email string) (*models.Email, error) {
	u := &models.Email{}
	if err := r.store.db.QueryRow(
		ctx,
		"SELECT id, email, confirmed FROM emails WHERE email=$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Confirmed,
	); err != nil {
		return nil, err
	}
	return u, nil
}
