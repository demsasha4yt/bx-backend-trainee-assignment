package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestEmailsRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("emails")

	s := sqlstore.New(db)
	m := models.TestEmail(t)

	assert.NoError(t, s.Emails().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, int64(0))

	m2 := models.TestEmail(t)
	assert.Error(t, s.Emails().Create(context.Background(), m2))
	assert.Equal(t, m2.ID, int64(0))
}

func TestEmailsRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("emails", "ads")

	s := sqlstore.New(db)
	m := models.TestEmail(t)

	assert.NoError(t, s.Emails().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, int64(0))

	found, err := s.Emails().FindByEmail(context.Background(), m.Email)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	found, err = s.Emails().FindByEmail(context.Background(), "abrakadabra@hhhhh.ru")
	assert.Error(t, err)
	assert.Nil(t, found)
}
