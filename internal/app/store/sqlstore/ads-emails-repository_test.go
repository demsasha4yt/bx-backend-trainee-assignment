package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestAdsEmailsRepository_FindByIds(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("ads", "emails")

	s := sqlstore.New(db)
	m := models.TestAd(t)
	u := models.TestEmail(t)
	assert.NoError(t, s.Emails().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, int64(0))

	m.Emails = append(m.Emails, u)

	assert.NoError(t, s.Ads().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, int64(0))

	found, err := s.AdsEmails().FindByIds(context.Background(), m.ID, u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

func TestAdsEmailsRepository_UpdateConfirmed(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("ads", "emails")

	s := sqlstore.New(db)
	m := models.TestAd(t)
	u := models.TestEmail(t)
	assert.NoError(t, s.Emails().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, int64(0))

	m.Emails = append(m.Emails, u)

	assert.NoError(t, s.Ads().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, int64(0))

	found, err := s.AdsEmails().FindByIds(context.Background(), m.ID, u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	assert.NoError(t, s.AdsEmails().UpdateConfirmed(context.Background(), found, true))
}

func TestAdsEmailsRepositoryDelete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("ads", "emails")

	s := sqlstore.New(db)
	m := models.TestAd(t)
	u := models.TestEmail(t)
	assert.NoError(t, s.Emails().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, int64(0))

	m.Emails = append(m.Emails, u)

	assert.NoError(t, s.Ads().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, int64(0))

	found, err := s.AdsEmails().FindByIds(context.Background(), m.ID, u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	assert.NoError(t, s.AdsEmails().Delete(context.Background(), found))

	found, err = s.AdsEmails().FindByIds(context.Background(), m.ID, u.ID)
	assert.Error(t, err)
	assert.Nil(t, found)
}
