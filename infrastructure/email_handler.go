package infrastructure

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"net/smtp"
)

type EmailMessenger struct {
	ce *domain.ConfigEmail
}

func NewEmailMessenger(ce *domain.ConfigEmail) usecases.Emailer {
	return &EmailMessenger{ce: ce}
}

func (em *EmailMessenger) SendMessage(to, subject, message string) error {
	mail := email.NewEmail()
	mail.From = em.ce.From
	mail.To = []string{to}
	mail.Subject = subject
	mail.HTML = []byte(message)
	return mail.Send(fmt.Sprintf("%s:%d", em.ce.Host, em.ce.Port), smtp.PlainAuth("", em.ce.Username, em.ce.Password, em.ce.Host))
	return nil
}
