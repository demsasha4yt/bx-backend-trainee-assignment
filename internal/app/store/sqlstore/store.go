package sqlstore

import (
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store ...
type Store struct {
	db                  *pgxpool.Pool
	adsRepository       *AdsRepository
	emailsRepository    *EmailsRepository
	adsEmailsRepository *AdsEmailsRepository
}

// New creates new Store
func New(db *pgxpool.Pool) store.Store {
	return &Store{
		db: db,
	}
}

// Ads returns AdsRepository
func (s *Store) Ads() store.AdsRepository {
	if s.adsRepository != nil {
		return s.adsRepository
	}
	s.adsRepository = &AdsRepository{
		store: s,
	}
	return s.adsRepository
}

// Emails returns EmailsRepository
func (s *Store) Emails() store.EmailsRepository {
	if s.emailsRepository != nil {
		return s.emailsRepository
	}
	s.emailsRepository = &EmailsRepository{
		store: s,
	}
	return s.emailsRepository
}

// AdsEmails returns AdsEmailsRepository
func (s *Store) AdsEmails() store.AdsEmailsRepository {
	if s.adsEmailsRepository != nil {
		return s.adsEmailsRepository
	}
	s.adsEmailsRepository = &AdsEmailsRepository{
		store: s,
	}
	return s.adsEmailsRepository
}
