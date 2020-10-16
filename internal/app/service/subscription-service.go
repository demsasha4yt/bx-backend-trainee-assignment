package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/emailsender"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// SubscriptionService service
type subscriptionService struct {
	service     *service
	emailSender emailsender.EmailSender
	avitoAPI    avitoapi.AvitoAPI
}

// Subscribe to ad
// Returns message and error
func (s *subscriptionService) Subscribe(ctx context.Context, email string, link string) (interface{}, error) {
	e, err := s.service.Emails().FindByEmailOrCreate(ctx, email)
	if err != nil {
		return "", err
	}

	ad, err := models.NewAdFromLink(link)
	if err != nil {
		return "", err
	}

	e.GenerateTokens(ad.AvitoID) // Generate tokens for confirm email or unsubscribe
	ad.Emails = append(ad.Emails, e)

	if err := ad.GetInfo(s.avitoAPI); err != nil {
		return "", err
	}

	if err := s.service.Ads().SubscribeOrCreate(ctx, ad); err != nil {
		return "", err
	}

	if !e.Confirmed {
		return "Подтвердите Email", nil
	}

	return fmt.Sprintf("Подписка на объявление №%d успешно активирована", ad.AvitoID), nil
}

func (s *subscriptionService) ConfirmSubscribe(ctx context.Context,
	emailID, adID int64, token string) (interface{}, error) {
	ad, err := s.service.AdsEmails().FindByIds(ctx, adID, emailID)

	if err != nil {
		return "", err
	}

	if ad.ConfirmToken != token {
		return "", errors.New("Unknown token")
	}

	if ad.Confirmed == true {
		return "Вы уже подтвердили Email", nil
	}

	err = s.service.AdsEmails().UpdateConfirmed(ctx, ad, true)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Вы успешно подтвердили email для рассылки №%d", ad.AdID), nil
}

func (s *subscriptionService) Unsubscribe(ctx context.Context,
	emailID, adID int64, token string) (interface{}, error) {
	ad, err := s.service.AdsEmails().FindByIds(ctx, adID, emailID)

	if err != nil {
		return "", err
	}

	if ad.UnsubscribeToken != token {
		return "", errors.New("Unknown token")
	}

	err = s.service.AdsEmails().Delete(ctx, ad)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Вы успешно отписались от рассылки №%d", ad.AdID), nil
}
