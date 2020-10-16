package emailsender

import (
	"fmt"
)

// GetConfirmationTemplate ...
func (s *emailSender) GetConfirmationTemplate(avitoID int64) (string, string) {
	subject := fmt.Sprintf("Подтвердите электронную почту для объявления №%d", avitoID)
	body := "Ссылка на подтверждение:"
	return subject, body
}

// GetConfirmationTemplate ...
func (s *emailSender) GetConfirmedTemplate(avitoID int64) (string, string) {
	subject := fmt.Sprintf("Подписка на объявление №%d активирована", avitoID)
	body := "Ссылка на подтверждение:"
	return subject, body
}

// GetPriceChangedTemplate ...
func (s *emailSender) GetPriceChangedTemplate(avitoID int64) (string, string) {
	subject := fmt.Sprintf("Стоимость объявления №%d изменилась.", avitoID)
	body := "Стоимость объявления изменилась"
	return subject, body
}

// GetConfirmationTemplate ...
func (s *emailSender) GetAdNotActualTemplate(avitoID int64) (string, string) {
	subject := fmt.Sprintf("Объявление №%d удалено пользователем", avitoID)
	body := "Объявление больше не актуально"
	return subject, body
}
