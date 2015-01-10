package domain

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type UserRepo interface {
	Store(user *User) error
	Update(user *User) error
	FindById(id string) (*User, error)
	FindByIdInt(id int64) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAdmin() ([]User, error)
	UpdatePassword(userId, password string) error
	GetAll() ([]User, error)
	Delete(id string) error
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Hash     string `json:"hash"`
	Password string `json:"-"`
	Bio      string `json:"bio"`
	Token    string `json:"token"`
	Role     Role   `json:"role"`
}

func (u *User) IsAdmin() bool {
	return u.Role == Admin
}

func (u *User) GenerateToken() {
	k := make([]byte, 32)
	io.ReadFull(rand.Reader, k)
	u.Token = base64.StdEncoding.EncodeToString(k)
}

func (u *User) HashEmail() {
	hasher := md5.New()
	hasher.Write([]byte(u.Email))
	u.Hash = hex.EncodeToString(hasher.Sum(nil))
}

func NewAdminUser(name, email, password string) (*User, error) {
	return NewUser(name, email, password, Admin)
}

func NewUser(name, email, password string, role Role) (*User, error) {
	if bArr, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err == nil {
		usr := &User{Name: name, Email: email, Password: string(bArr), Role: role}
		usr.GenerateToken()
		usr.HashEmail()
		return usr, nil
	} else {
		return nil, err
	}
}
