package domain

import (
	"strings"
	"time"
)

type CategoryRepo interface {
	Store(category *Category) error
	FindById(id int) (*Category, error)
	FindByTitle(title string) (*Category, error)
	GetAll() ([]*Category, error)
	DeleteById(id int) error
}

type Category struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Slug    string    `json:"slug"`
	Count   int       `json:"count"`
	Created time.Time `json:"create"`
}

func (c *Category) GenerateSlug() {
	if c.Slug == "" {
		slug := c.Title
		slug = strings.Replace(slug, " ", SLUG_SPACER, -1)
		slug = strings.ToLower(slug)
		c.Slug = slug
	}
}
