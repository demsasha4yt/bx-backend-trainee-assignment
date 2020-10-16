package avitoapi

import "math/rand"

// AvitoAPI interface
type AvitoAPI interface {
	GetInfo(avitoID int64) (*Response, error)
}

type avitoAPI struct {
	conf *Config
}

// Response ...
type Response struct {
	Price int
}

// New returns new AvitoAPI
func New(conf *Config) AvitoAPI {
	return &avitoAPI{
		conf: conf,
	}
}

// GetInfo returns info about ad
func (s *avitoAPI) GetInfo(avitoID int64) (*Response, error) {
	return &Response{
		Price: rand.Intn(10000),
	}, nil
}
