package service

import (
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/emailsender"
)

type priceTrackerService struct {
	service     Service
	emailSender emailsender.EmailSender
	avitoAPI    avitoapi.AvitoAPI
}

func (s *priceTrackerService) CheckAdsTask() error {
	return nil
}
