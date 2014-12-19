package usecases

type Emailer interface {
	SendMessage(to, subject, message string) error
}
