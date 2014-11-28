package domain

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	STRLEN = 15
)

type UserRepo interface {
	Store(user *User) error
	FindById(id string) (*User, error)
	FindByIdInt(id int64) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAdmin() (*[]User, error)
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
}

func (u *User) IsAdmin() bool {
	return u.Role == Admin
}

func NewAdminUser(name, email, password string) (*User, error) {
	if bArr, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err == nil {
		usr := &User{Name: name, Email: email, Password: string(bArr), Role: Admin}
		return usr, nil
	} else {
		return nil, err
	}
}
