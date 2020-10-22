package avitoapi_test

import (
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	"github.com/stretchr/testify/assert"
)

func TestAvitoAPI_GetInfo(t *testing.T) {
	avitoAPI := avitoapi.TestAvitoAPI(t)
	res, err := avitoAPI.GetInfo(int64(100301))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	res, err = avitoAPI.GetInfo(int64(2))
	assert.Error(t, err)
	assert.Nil(t, res)
}
