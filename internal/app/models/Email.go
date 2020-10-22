package models

import (
	"encoding/base64"
	"encoding/json"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

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

// Validate validates struct
func (m *Email) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Email, validation.Required, is.Email),
	)
}

// GenerateTokens genererates unsubsribe and confirmation tokens
func (m *Email) GenerateTokens(adID int64) error {
	if err := m.generateConfirmToken(adID); err != nil {
		return err
	}
	if err := m.generateUnsubscribeToken(adID); err != nil {
		return err
	}
	return nil
}

// generateConfirmToken generates confirm token
func (m *Email) generateConfirmToken(adID int64) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(strconv.FormatInt(adID, 10)+m.Email+"_confirm"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	m.ConfirmToken = base64.StdEncoding.EncodeToString(hash)
	return nil
}

// generateUnsubscribeToken generates unsubcribe token
func (m *Email) generateUnsubscribeToken(adID int64) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(strconv.FormatInt(adID, 10)+m.Email+"_unsibscribe"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	m.UnsubscribeToken = base64.StdEncoding.EncodeToString(hash)
	return nil
}
