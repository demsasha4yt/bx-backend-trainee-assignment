package models

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

// Validate validates struct
func (m *Ad) Validate() error {
	return nil
}

// SetAvitoIDFromLink parses link and put it to AvitoID field
func (m *Ad) SetAvitoIDFromLink(link string) error {
	return nil
}
