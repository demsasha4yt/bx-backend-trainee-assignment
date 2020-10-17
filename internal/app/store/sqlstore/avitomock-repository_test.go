package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/stretchr/testify/assert"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store/sqlstore"
)

func TestAvitoMockRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("avito_mockapi")

	s := sqlstore.New(db)
	u := models.TestAvitoMock(t)

	assert.NoError(t, s.AvitoMock().Create(context.Background(), u))
	assert.NotEqual(t, 0, u.ID)
}

func TestAvitoMockRepository_FindByAvitoID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("avito_mockapi")

	s := sqlstore.New(db)
	u := models.TestAvitoMock(t)

	assert.NoError(t, s.AvitoMock().Create(context.Background(), u))
	assert.NotEqual(t, 0, u.ID)

	found, err := s.AvitoMock().FindByAvitoID(context.Background(), u.AvitoID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	found, err = s.AvitoMock().FindByAvitoID(context.Background(), int64(1200044))
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestAvitoMockRepository_UpdatePrices(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("avito_mockapi")

	s := sqlstore.New(db)

	assert.NoError(t, s.AvitoMock().UpdatePrices(context.Background()))
}

func TestAvitoMockRepository_SetDeleted(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("avito_mockapi")

	s := sqlstore.New(db)

	assert.NoError(t, s.AvitoMock().SetDeleted(context.Background()))
}

func TestAvitoMockRepository_UpdatePrice(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("avito_mockapi")

	s := sqlstore.New(db)
	u := models.TestAvitoMock(t)

	assert.NoError(t, s.AvitoMock().Create(context.Background(), u))
	assert.NotEqual(t, 0, u.ID)

	assert.NoError(t, s.AvitoMock().UpdatePrice(context.Background(), u.AvitoID, 10000))

	found, err := s.AvitoMock().FindByAvitoID(context.Background(), u.AvitoID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	assert.Equal(t, found.Price, 10000)
}

func TestAvitoMockRepository_SetDeletedOne(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("avito_mockapi")

	s := sqlstore.New(db)
	u := models.TestAvitoMock(t)

	assert.NoError(t, s.AvitoMock().Create(context.Background(), u))
	assert.NotEqual(t, 0, u.ID)

	assert.NoError(t, s.AvitoMock().SetDeletedOne(context.Background(), u.AvitoID))

	found, err := s.AvitoMock().FindByAvitoID(context.Background(), u.AvitoID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	assert.Equal(t, found.Deleted, true)
}
