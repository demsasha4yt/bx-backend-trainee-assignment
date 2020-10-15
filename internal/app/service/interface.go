package service

import "context"

// Interface ...
type Interface interface {
	Subscription() SubscriptionServiceInterface
}

// SubscriptionServiceInterface ...
type SubscriptionServiceInterface interface {
	Subscribe(context.Context, string, string) (interface{}, error)
}
