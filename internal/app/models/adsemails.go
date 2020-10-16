package models

// AdsEmails model
type AdsEmails struct {
	ID               int64  `json:"id"`
	AdID             int64  `json:"ad_id"`
	EmailID          int64  `json:"email_id"`
	Confirmed        bool   `json:"confirmed"`
	ConfirmToken     string `json:"confirm_token"`
	UnsubscribeToken string `json:"unsubscribe_token"`
}

// Validate validates struct
func (m *AdsEmails) Validate() error {
	return nil
}
