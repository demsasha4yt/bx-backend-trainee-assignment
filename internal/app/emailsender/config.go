package emailsender

// EmailConfig ...
type EmailConfig struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	ServerHost string `json:"server_host"`
	ServerPort string `json:"server_port"`
	SenderAddr string `json:"sender_addr"`
}
