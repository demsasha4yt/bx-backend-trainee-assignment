package emailsender

import "net/smtp"

// EmailSender ...
type EmailSender interface {
	Send(to []string, body []byte) error
}

type emailSender struct {
	conf Config
	send func(string, smtp.Auth, string, []string, []byte) error
}

// NewEmailSender ...
func NewEmailSender(conf Config) EmailSender {
	return &emailSender{conf, smtp.SendMail}
}

func (e *emailSender) Send(to []string, body []byte) error {
	addr := e.conf.ServerHost + ":" + e.conf.ServerPort
	auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)
	return e.send(addr, auth, e.conf.SenderAddr, to, body)
}
