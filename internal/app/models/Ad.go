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
	PriceHistory []int    `json:"price_history,omitepty"`
}
