package usecases

import (
	"github.com/kendellfab/publish/domain"
)

type RepoManager struct {
	CommentRepo  domain.CommentRepo
	ContactRepo  domain.ContactRepo
	UserRepo     domain.UserRepo
	CategoryRepo domain.CategoryRepo
	PostRepo     domain.PostRepo
	PageRepo     domain.PageRepo
	PayloadRepo  domain.PayloadRepo
}
