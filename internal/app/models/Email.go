package models

// Email model ...
type Email struct {
	ID        int64  `json:"id"`
	Email     string `json:"email,omitempty"`
	Confirmed bool   `json:"confirmed,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
