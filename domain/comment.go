package domain

import (
	"time"
)

type CommentRepo interface {
	Store(comment *Comment) error
	FindById(id int) (*Comment, error)
	FindByPage(page string) (*[]Comment, error)
	FindUnapprovedComments() (*[]Comment, error)
	ApproveComment(id int) error
}

type Comment struct {
	Id       int       `json:"id"`
	Page     string    `json:"page"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Date     time.Time `json:"date"`
	Content  string    `json:"content"`
	Approved bool      `json:"approved"`
}
