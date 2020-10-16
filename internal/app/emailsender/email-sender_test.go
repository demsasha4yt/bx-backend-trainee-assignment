package emailsender_test

import (
	"testing"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/emailsender"
	"github.com/stretchr/testify/assert"
)

func TestEmailSender_Send(t *testing.T) {
	c := emailsender.Config{
		ServerHost: "smtp.yandex.ru",
		ServerPort: "465",
		Username:   "info@bharrold.ru",
		Password:   "contact me if u need password",
		SenderAddr: "info@bharrold.ru",
	}
	es := emailsender.NewEmailSender(c)
	assert.NoError(t, es.Send("demsasha4yt@yandex.ru",
		"Тестовое сообщение",
		"Привет2"))
}
