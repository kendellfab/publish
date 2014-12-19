package infrastructure

import (
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
)

type EmailMessenger struct {
	ce *domain.ConfigEmail
}

func NewEmailMessenger(ce *domain.ConfigEmail) usecases.Emailer {
	return &EmailMessenger{ce: ce}
}

func (em *EmailMessenger) SendMessage(to, subject, message string) error {
	return nil
}
