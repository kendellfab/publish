package domain

type ContactRepo interface {
	Store(contact *ContactForm) error
	GetAll() (*[]ContactForm, error)
	Delete(id int) error
}

type ContactForm struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
