package models_test

import (
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestAdsEmails_Validate(t *testing.T) {
	u := models.AdsEmails{
		AdID:    1,
		EmailID: 1,
	}
	assert.NoError(t, u.Validate())
	testCases := []struct {
		name    string
		u       func() *models.AdsEmails
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
