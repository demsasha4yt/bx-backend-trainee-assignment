package avitomock

// Config for Avito API Mock
type Config struct {
	BindAddr    string `json:"bind_addr"`
	LogLevel    string `json:"log_level"`
	DatabaseURL string `json:"database_url"`
	Delay       int    `json:"delay"`
	AvitoKey    string `json:"key"`
}

// NewConfig ..
func NewConfig() *Config {
	return &Config{
		BindAddr: ":9000",
		LogLevel: "debug",
	}
}
