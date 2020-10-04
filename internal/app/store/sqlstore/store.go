package sqlstore

import (
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store ...
type Store struct {
	db               *pgxpool.Pool
	adsRepository    *AdsRepository
	emailsRepository *EmailsRepository
}

// New creates new Store
func New(db *pgxpool.Pool) store.Store {
	return &Store{
		db: db,
	}
}
