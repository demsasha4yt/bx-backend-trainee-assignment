package emailsender

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"
)

// EmailTemplate struct....
type EmailTemplate struct {
	Subject string
	Body    string
}

const (
	confirmationTemplateFile = "./templates/confirmation.html"
	confirmedTemplateFile    = "./templates/confirmed.html"
	notActualTemplateFile    = "./templates/notactual.html"
	priceChangedTemplateFile = "./templates/pricechanged.html"
)

type emailData struct {
	AvitoID          int64
	Link             string
	CurrentPrice     int
	PriceHistory     []int
	ConfirmationLink string
	UnsubscribeLink  string
}

// GetConfirmationTemplate ...
func GetConfirmationTemplate(ad *models.Ad, email *models.Email) (*EmailTemplate, error) {
	d := &emailData{
		AvitoID:          ad.AvitoID,
		Link:             ad.Link,
		ConfirmationLink: "",
	}
	body, err := parseTemplate(confirmationTemplateFile, d)
	if err != nil {
		return nil, err
	}

	subject := fmt.Sprintf("Подтвердите электронную почту для объявления №%d", ad.AvitoID)

	return &EmailTemplate{subject, body}, nil
}

// GetConfirmedTemplate ...
func GetConfirmedTemplate(ad *models.Ad, email *models.Email) (*EmailTemplate, error) {
	d := &emailData{
		AvitoID:         ad.AvitoID,
		Link:            ad.Link,
		UnsubscribeLink: "",
	}
	body, err := parseTemplate(confirmedTemplateFile, d)
	if err != nil {
		return nil, err
	}

	subject := fmt.Sprintf("Подписка на объявление №%d активирована", ad.AvitoID)

	return &EmailTemplate{subject, body}, nil
}

// GetPriceChangedTemplate ...
func GetPriceChangedTemplate(ad *models.Ad, email *models.Email) (*EmailTemplate, error) {
	d := &emailData{
		AvitoID:         ad.AvitoID,
		Link:            ad.Link,
		UnsubscribeLink: "",
	}
	body, err := parseTemplate(priceChangedTemplateFile, d)
	if err != nil {
		return nil, err
	}

	subject := fmt.Sprintf("Стоимость объявления №%d изменилась.", ad.AvitoID)

	return &EmailTemplate{subject, body}, nil
}

// GetAdNotActualTemplate ...
func GetAdNotActualTemplate(ad *models.Ad, email *models.Email) (*EmailTemplate, error) {
	d := &emailData{
		AvitoID: ad.AvitoID,
		Link:    ad.Link,
	}
	body, err := parseTemplate(notActualTemplateFile, d)
	if err != nil {
		return nil, err
	}

	subject := fmt.Sprintf("Объявление №%d удалено пользователем", ad.AvitoID)

	return &EmailTemplate{subject, body}, nil
}

// parseTemplate parse email body from html template
func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
