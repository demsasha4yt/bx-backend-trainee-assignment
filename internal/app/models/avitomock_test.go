package models_test

import (
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func Test_AvitoMock_NewAvitoMockFromID(t *testing.T) {
	id := int64(64)
	u := models.NewAvitoMockFromID(id)
	assert.NotNil(t, u)
	assert.Equal(t, u.AvitoID, int64(64))
	assert.NotEqual(t, u.Price, 0)
}

func TestAvitoMock_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.AvitoMock
		isValid bool
	}{
		{
			name: "valid",
			u: func() *models.AvitoMock {
				return &models.AvitoMock{
					AvitoID: 1,
					Price:   1,
				}
			},
			isValid: true,
		},
		{
			name: "not-valid-avitoid",
			u: func() *models.AvitoMock {
				return &models.AvitoMock{
					AvitoID: 0,
					Price:   1,
				}
			},
			isValid: false,
		},
		{
			name: "not-valid-price",
			u: func() *models.AvitoMock {
				return &models.AvitoMock{
					AvitoID: 1,
					Price:   0,
				}
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestAvitoMock_GenerateRandomPrice(t *testing.T) {
	u := models.TestAvitoMock(t)
	u.Price = 0
	u.GenerateRandomPrice()
	assert.NotEqual(t, u.Price, 0)
}
