package service

import (
	"context"
)

// SubscriptionService service
type subscriptionService struct {
	service *service
}

// Subscribe to ad
// Returns message and error
func (s *subscriptionService) Subscribe(ctx context.Context, email string, link string) (interface{}, error) {
	e, err := s.service.Emails().FindByEmailOrCreate(ctx, email)
	if err != nil {
		return "", err
	}
	e.GenerateTokens(int64(10))
	return e, nil
}

func (s *subscriptionService) ConfirmSubscribe(ctx context.Context,
	emailID, adID int, token string) (interface{}, error) {
	return "ConfirmSubscribe", nil
}

func (s *subscriptionService) Unsubscribe(ctx context.Context,
	emailID, adID int, token string) (interface{}, error) {
	return "Unsubscribe", nil
}
