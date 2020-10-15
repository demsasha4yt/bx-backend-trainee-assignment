package models

import "testing"

// TestEmail returns testing Email model
func TestEmail(t *testing.T) *Email {
	t.Helper()

	return &Email{
		Email: "test@yandex.ru",
	}
}

// TestAd returns testing Ad model
func TestAd(t *testing.T) *Ad {
	t.Helper()
	return &Ad{
		AvitoID:      100,
		Actual:       true,
		CurrentPrice: 2000,
		Emails:       make([]*Email, 0),
	}
}
