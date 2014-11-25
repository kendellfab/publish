package domain

const (
	STRLEN     = 15
	ROLE_ADMIN = "role_admin"
	ROLE_USER  = "role_user"
)

type UserRepo interface {
	Store(user *User) error
	FindById(id string) (*User, error)
	FindByIdInt(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAdmin() (*[]User, error)
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

func (u *User) IsAdmin() bool {
	return u.Role == ROLE_ADMIN
}
