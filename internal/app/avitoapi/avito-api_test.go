package avitoapi_test

import (
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	"github.com/stretchr/testify/assert"
)

func TestAvitoAPI_GetInfo(t *testing.T) {
	const AvitoID = int64(1987554999)
	_, err := avitoapi.GetInfo(AvitoID)
	assert.NoError(t, err)
	t.Fatal()
}
