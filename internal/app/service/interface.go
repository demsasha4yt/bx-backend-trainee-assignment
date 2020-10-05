package service

// Interface ...
type Interface interface {
	Subscription() SubscriptionServiceInterface
}

// SubscriptionServiceInterface ...
type SubscriptionServiceInterface interface {
	Subscribe(string, string) (string, error)
}
