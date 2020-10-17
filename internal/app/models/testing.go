package models

import "testing"

// TestEmail returns testing Email model
func TestEmail(t *testing.T) *Email {
	t.Helper()

	u := &Email{
		Email: "test@yandex.ru",
	}
	u.GenerateTokens(int64(0))
	return u
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

// TestAvitoMock returns testing AvitoMock
func TestAvitoMock(t *testing.T) *AvitoMock {
	t.Helper()

	return &AvitoMock{
		AvitoID: 10000,
		Price:   10000,
	}
}
