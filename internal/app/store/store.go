package store

// Store ...
type Store interface {
	Ads() AdsRepository
	Emails() EmailsRepository
	AdsEmails() AdsEmailsRepository
}
