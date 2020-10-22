package models_test

import (
	"fmt"
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAd_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.Ad
		isValid bool
	}{
		{
			name: "valid",
			u: func() *models.Ad {
				return &models.Ad{
					CurrentPrice: 1,
					AvitoID:      1,
				}
			},
			isValid: true,
		},
		{
			name: "valid-with-emails",
			u: func() *models.Ad {
				return &models.Ad{
					CurrentPrice: 1,
					AvitoID:      1,
					Emails: []*models.Email{
						&models.Email{
							Email: "test@gmail.com",
						},
					},
				}
			},
			isValid: true,
		},
		{
			name: "not-valid-with-emails",
			u: func() *models.Ad {
				return &models.Ad{
					CurrentPrice: 1,
					AvitoID:      1,
					Emails: []*models.Email{
						&models.Email{
							Email: "testsdafsadfasdf",
						},
					},
				}
			},
			isValid: false,
		},
		{
			name: "not-valid-avitoid",
			u: func() *models.Ad {
				return &models.Ad{
					CurrentPrice: 1,
					AvitoID:      0,
				}
			},
			isValid: false,
		},
		{
			name: "not-valid-current-price",
			u: func() *models.Ad {
				return &models.Ad{
					CurrentPrice: 0,
					AvitoID:      1,
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

type avitoAPIMock struct {
	mock.Mock
}

func (m *avitoAPIMock) GetInfo(avitoID int64) (*avitoapi.Response, error) {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", avitoID)
	args := m.Called(avitoID)
	return &avitoapi.Response{Price: int(avitoID)}, args.Error(0)
}

func TestAd_GetInfo(t *testing.T) {
	mock := new(avitoAPIMock)

	mock.On("GetInfo", int64(1000)).Return(nil)
	mock.On("GetInfo", int64(1)).Return(avitoapi.ErrNotFound)
	mock.On("GetInfo", int64(2)).Return(avitoapi.ErrNotOK)
	mock.On("GetInfo", int64(3)).Return(nil)
	ad := &models.Ad{
		CurrentPrice: 1,
		AvitoID:      int64(1000),
	}
	assert.NoError(t, ad.GetInfo(mock))
	assert.Equal(t, int(ad.AvitoID), ad.CurrentPrice)
	ad.AvitoID = int64(1)
	assert.EqualError(t, ad.GetInfo(mock), avitoapi.ErrNotFound.Error())
	ad.AvitoID = int64(2)
	assert.EqualError(t, ad.GetInfo(mock), avitoapi.ErrNotOK.Error())
	ad.AvitoID = int64(3)
	assert.NoError(t, ad.GetInfo(mock))
	assert.Equal(t, int(ad.AvitoID), ad.CurrentPrice)
	mock.AssertExpectations(t)
}
