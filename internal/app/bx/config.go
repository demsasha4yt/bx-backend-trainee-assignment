package bx

import (
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/emailsender"
)

// Config ...
type Config struct {
	BindAddr       string             `json:"bind_addr"`
	LogLevel       string             `json:"log_level"`
	DatabaseURL    string             `json:"database_url"`
	SMTPConfig     emailsender.Config `json:"email_sender"`
	AvitoAPIConfig *avitoapi.Config   `json:"avito_api"`
}

// NewConfig ..
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
		LogLevel: "debug",
	}
}
