package domain

import (
	"time"
)

type ViewRepo interface {
	Store(v *View) error
	GetType(t TargetType) ([]*View, error)
	GetTypeTarget(t TargetType, target string) ([]*View, error)
}

type TargetType int

const (
	TypeUnkown TargetType = iota
	TypePage
	TypePost
)

func (t TargetType) String() string {
	switch t {
	case 1:
		return "Page"
	case 2:
		return "Post"
	default:
		return "Unkown"
	}
}

type View struct {
	Id         int        `json:"id"`
	From       string     `json:"from"`
	When       time.Time  `json:"when"`
	TargetType TargetType `json:"targetType"`
	Target     string     `json:"target"`
}
