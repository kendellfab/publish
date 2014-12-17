package usecases

type Emailer interface {
	SendMessage(to, message string) error
}
