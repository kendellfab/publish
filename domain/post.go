package domain

import (
	"strings"
	"time"
)

type PostRepo interface {
	Store(post *Post) error
	Update(post *Post) error
	FindById(id int64) (*Post, error)
	FindByIdString(id string) (*Post, error)
	FindBySlug(slug string) (*Post, error)
	FindByCategory(category *Category) (*[]Post, error)
	FindAll() (*[]Post, error)
	Delete(id int) error
	Publish(id int) error
	UnPublish(id int) error
}

type Post struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Author      User      `json:"author"`
	AuthorId    int64     `json:"authorId"`
	Created     time.Time `json:"created"`
	Content     string    `json:"content"`
	ContentType string    `json:"content_type"`
	Published   bool      `json:"published"`
	Tags        []string  `json:"tags"`
	Category    Category  `json:"category"`
	Comments    []Comment `json:"comments"`
}

func (p *Post) GenerateSlug() {
	if p.Slug == "" {
		slug := p.Title
		slug = strings.Replace(slug, " ", SLUG_SPACER, -1)
		slug = strings.ToLower(slug)
		p.Slug = slug
	}
}
