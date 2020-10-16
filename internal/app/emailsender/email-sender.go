package emailsender

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

// EmailSender ...
type EmailSender interface {
	Send(to string, subject string, body string) error
}

type emailSender struct {
	conf Config
}

// NewEmailSender ...
func NewEmailSender(conf Config) EmailSender {
	fmt.Println("CONFIGURE SMTP SERVER:")
	fmt.Println("\tSMTP_HOST:", conf.ServerHost)
	fmt.Println("\tSMTP_PORT:", conf.ServerPort)
	fmt.Println("\tSMTP_USER:", conf.Username)
	fmt.Println("\tSMTP_PASSWORD:", conf.Password)
	fmt.Println("\tSMTP_SENDER_ADDR:", conf.SenderAddr)
	return &emailSender{conf}
}

func (e *emailSender) Send(to string, subject string, body string) error {

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = e.conf.SenderAddr
	headers["To"] = to
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	addr := e.conf.ServerHost + ":" + e.conf.ServerPort

	auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         e.conf.ServerHost,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, e.conf.ServerHost)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(e.conf.SenderAddr); err != nil {
		return err
	}
	if err = c.Rcpt(to); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil
}
