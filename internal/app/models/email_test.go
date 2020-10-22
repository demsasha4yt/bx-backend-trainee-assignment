package models_test

import (
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestEmail_NewEmailFromByte(t *testing.T) {
	str := `{"email": "test@gmail.com"}`
	u, err := models.NewEmailFromByte([]byte(str))
	assert.NoError(t, err)
	assert.NotNil(t, u)

	str = `{"email"": "test@gmail.com"}`
	u, err = models.NewEmailFromByte([]byte(str))
	assert.Error(t, err)
	assert.Nil(t, u)
}

func TestEmail_NewEmailSliceFromByte(t *testing.T) {
	str := `[{"email": "test@gmail.com"}, {"email": "test@gmail.com"}]`
	u, err := models.NewEmailSliceFromByte([]byte(str))
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Len(t, u, 2)

	str = `[{"email": "test@gmail.com"]]`
	u, err = models.NewEmailSliceFromByte([]byte(str))
	assert.Error(t, err)
	assert.Nil(t, u)
}

func TestEmail_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.Email
		isValid bool
	}{
		{
			name: "valid",
			u: func() *models.Email {
				return &models.Email{
					Email: "test@gmail.com",
				}
			},
			isValid: true,
		},
		{
			name: "not-valid",
			u: func() *models.Email {
				return &models.Email{
					Email: "testddasdfasdfsad",
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

func TestEmail_GenerateTokens(t *testing.T) {
	u := models.TestEmail(t)

	assert.NoError(t, u.GenerateTokens(int64(5000)))
}
