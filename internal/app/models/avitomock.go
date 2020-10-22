package models

import (
	"math/rand"

	validation "github.com/go-ozzo/ozzo-validation"
)

// AvitoMock model
type AvitoMock struct {
	ID        int    `json:"id"`
	AvitoID   int64  `json:"avito_id,omitempty"`
	Price     int    `json:"price,omitempty"`
	Deleted   bool   `json:"deleted,omitempty"`
	CreatedAt string `json:"-"`
	ChangedAt string `json:"-"`
}

// NewAvitoMockFromID returns new AvitoMock model with random price
func NewAvitoMockFromID(avitoID int64) *AvitoMock {
	u := &AvitoMock{
		AvitoID: avitoID,
		Deleted: false,
	}
	u.GenerateRandomPrice()
	return u
}

// Validate validates AvitoMock Model
func (m *AvitoMock) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.AvitoID, validation.Required, validation.Min(1)),
		validation.Field(&m.Price, validation.Required, validation.Min(1)),
	)
}

// GenerateRandomPrice generates random price for AvitoMock
func (m *AvitoMock) GenerateRandomPrice() {
	m.Price = rand.Intn(10000)
}
