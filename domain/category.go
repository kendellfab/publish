package domain

import (
	"time"
)

type CategoryRepo interface {
	Store(category *Category) error
	FindById(id int) (*Category, error)
	FindByTitle(title string) (*Category, error)
	GetAll() (*[]Category, error)
	GetCategoryPostCount() (*[]CategoryCount, error)
	DeleteById(id int) error
}

type Category struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"create"`
}

type CategoryCount struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}
