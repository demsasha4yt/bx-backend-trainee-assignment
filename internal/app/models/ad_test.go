package models_test

import (
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestAd_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.Ad
		isValid bool
	}{}

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

func TestAd_NewAdFromLink(t *testing.T) {
	testCases := []struct {
		link    string
		avitoID int64
		isValid bool
	}{
		{
			link:    "https://www.avito.ru/moskva/sport_i_otdyh/joma_indoor_dribling_2013585850",
			avitoID: int64(2013585850),
			isValid: true,
		},
		{
			link:    "https://www.avito.ru/moskva/sport_i_otdyh/joma_indoor_dribling",
			avitoID: 0,
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.link, func(t *testing.T) {
			u, err := models.NewAdFromLink(tc.link)
			if tc.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tc.avitoID, u.AvitoID)
			} else {
				assert.Error(t, err)
				assert.Nil(t, u)
			}
		})
	}
}
