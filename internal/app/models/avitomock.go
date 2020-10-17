package models

import "math/rand"

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
	u := &AvitoMock{}
	u.GenerateRandomPrice()
	return u
}

// Validate validates AvitoMock Model
func (m *AvitoMock) Validate() error {
	return nil
}

// GenerateRandomPrice generates random price for AvitoMock
func (m *AvitoMock) GenerateRandomPrice() {
	m.Price = rand.Intn(10000)
}
