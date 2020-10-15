package models

import (
	"encoding/base64"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

// Email model ...
type Email struct {
	ID               int64  `json:"id"`
	Email            string `json:"email,omitempty"`
	Confirmed        bool   `json:"confirmed,omitempty"`
	ConfirmToken     string `json:"-"`
	UnsubscribeToken string `json:"-"`
	CreatedAt        string `json:"created_at,omitempty"`
}

// Validate validates struct
func (m *Email) Validate() error {
	return nil
}

// GenerateTokens genererates unsubsribe and confirmation tokens
func (m *Email) GenerateTokens(adID int) error {
	if err := m.generateConfirmToken(adID); err != nil {
		return err
	}
	if err := m.generateUnsubscribeToken(adID); err != nil {
		return err
	}
	return nil
}

// NewEmailFromByte creates struct from byte slice
func NewEmailFromByte(b []byte) (*Email, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Email{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewEmailSliceFromByte creates struct from byte slice
func NewEmailSliceFromByte(b []byte) ([]*Email, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Email = make([]*Email, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// generateConfirmToken generates confirm token
func (m *Email) generateConfirmToken(adID int) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(string(adID)+m.Email+"_confirm"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	m.ConfirmToken = base64.StdEncoding.EncodeToString(hash)
	return nil
}

// generateUnsubscribeToken generates unsubcribe token
func (m *Email) generateUnsubscribeToken(adID int) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(string(adID)+m.Email+"_unsibscribe"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	m.UnsubscribeToken = base64.StdEncoding.EncodeToString(hash)
	return nil
}
