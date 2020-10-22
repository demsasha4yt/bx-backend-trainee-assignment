package models

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Ad model
type Ad struct {
	ID           int64    `json:"id"`
	AvitoID      int64    `json:"avito_id,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	CheckedAt    string   `json:"checked_at,omitempty"`
	Actual       bool     `json:"actual,omitempty"`
	Emails       []*Email `json:"emails,omitempty"`
	CurrentPrice int      `json:"cur_price,omitempty"`
	PriceHistory []int    `json:"price_history,omitempty"`
}

// NewAdFromLink creates Ad from link
func NewAdFromLink(link string) (*Ad, error) {
	re := regexp.MustCompile(`([0-9]{1,})$`)
	avitoIDstring := re.FindString(link)
	if avitoIDstring == "" {
		return nil, fmt.Errorf("Can't find avitoID at the end of %s", link)
	}
	avitoID, err := strconv.ParseInt(avitoIDstring, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("Can't convert avitoID from %s", avitoIDstring)
	}
	u := &Ad{
		AvitoID: avitoID,
	}
	return u, nil
}

// GetInfo gets info about ad from avitoAPI
func (m *Ad) GetInfo(api avitoapi.AvitoAPI) error {
	resp, err := api.GetInfo(m.AvitoID)
	if err != nil {
		return err
	}
	m.CurrentPrice = resp.Price
	return nil
}

// Validate validates struct
func (m *Ad) Validate() error {
	fmt.Printf("%d %d\n", m.AvitoID, m.CurrentPrice)
	return validation.ValidateStruct(
		m,
		validation.Field(&m.AvitoID, validation.Required, validation.Min(1)),
		validation.Field(&m.CurrentPrice, validation.Required, validation.Min(1)),
		validation.Field(&m.Emails),
	)
}
