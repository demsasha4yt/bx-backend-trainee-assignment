package avitoapi

// Config for Avito API
type Config struct {
	URL string `json:"avito_url" env:"AVITO_ADDR"`
	Key string `json:"avito_key" env:"AVITO_KEY"`
}
