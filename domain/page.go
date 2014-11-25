package domain

import (
	"strings"
	"time"
)

type PageRepo interface {
	Store(page *Page) error
	FindById(id int) (*Page, error)
	FindBySlug(slug string) (*Page, error)
	FindAll() (*[]Page, error)
	Update(page *Page) error
	Publish(id int) error
	UnPublish(id int) error
	Delete(id int) error
	FindAllPublished() (*[]Page, error)
}

type Page struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Created   time.Time `json:"created"`
	Content   string    `json:"content"`
	Published bool      `json:"published"`
}

func (p *Page) GenerateSlug() {
	slug := p.Title
	slug = strings.Replace(slug, " ", SLUG_SPACER, -1)
	slug = strings.ToLower(slug)
	p.Slug = slug
}
