package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestAdsRepository_Create(t *testing.T) {
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
}
