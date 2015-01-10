package domain

type Role int

const (
	Unknown Role = iota
	Usr
	_
	_
	Author
	_
	_
	Admin
)

func (r Role) String() string {
	switch r {
	case Unknown:
		return "Unknown"
	case Usr:
		return "User"
	case Author:
		return "Author"
	case Admin:
		return "Admin"
	default:
		return ""
	}
}
