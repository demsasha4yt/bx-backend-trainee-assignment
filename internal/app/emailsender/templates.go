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

// GetConfirmationTemplate ...
func GetConfirmationTemplate(ad *models.Ad, email *models.Email) (*EmailTemplate, error) {
	subject := fmt.Sprintf("Подтвердите электронную почту для объявления №%d", ad.AvitoID)

	body, err := parseTemplate(confirmationTemplateFile, "")
	if err != nil {
		return nil, err
	}
	return &EmailTemplate{subject, body}, nil
}

// GetConfirmedTemplate ...
func GetConfirmedTemplate(ad *models.Ad, email *models.Email) (*EmailTemplate, error) {
	subject := fmt.Sprintf("Подписка на объявление №%d активирована", ad.AvitoID)

	body, err := parseTemplate(confirmedTemplateFile, "")
	if err != nil {
		return nil, err
	}
	return &EmailTemplate{subject, body}, nil
}

// GetPriceChangedTemplate ...
func GetPriceChangedTemplate(ad *models.Ad, email *models.Email) (*EmailTemplate, error) {
	subject := fmt.Sprintf("Стоимость объявления №%d изменилась.", ad.AvitoID)

	body, err := parseTemplate(priceChangedTemplateFile, "")
	if err != nil {
		return nil, err
	}
	return &EmailTemplate{subject, body}, nil
}

// GetAdNotActualTemplate ...
func GetAdNotActualTemplate(ad *models.Ad, email *models.Email) (*EmailTemplate, error) {
	subject := fmt.Sprintf("Объявление №%d удалено пользователем", ad.AvitoID)
	body, err := parseTemplate(notActualTemplateFile, "")
	if err != nil {
		return nil, err
	}
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
