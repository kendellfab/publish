package domain

import (
	"time"
)

type ResetRepo interface {
	Store(r *Reset) error
	FindByToken(token string) (*Reset, error)
	CleanExpired() error
}

type Reset struct {
	Id      int64
	UserId  int64
	Created time.Time
	Expires time.Time
	Token   string
}
