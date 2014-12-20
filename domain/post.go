package domain

import (
	"fmt"
	"strings"
	"time"
)

type PostRepo interface {
	Store(post *Post) error
	Update(post *Post) error
	FindById(id int64) (*Post, error)
	FindByIdString(id string) (*Post, error)
	FindBySlug(slug string) (*Post, error)
	FindByCategory(category *Category, offset, limit int) ([]*Post, error)
	FindAll() ([]*Post, error)
	FindPublished(offset, limit int) ([]*Post, error)
	FindByYearMonth(year, month string) ([]*Post, error)
	FindDashboard(offset, limit int) ([]*Post, error)
	Delete(id int) error
	Publish(id int64) error
	UnPublish(id int64) error
	PublishedCount() (int, error)
	PublishedCountCategory(catId int) (int, error)
	AddToSeries(id, seriesId string) error
	GetForSeries(seriesId string) ([]*Post, error)
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
	Category    *Category `json:"category"`
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

func (p Post) Description() string {
	index := strings.Index(p.Content, "<!--desc-->")
	if index != -1 {
		return p.Content[0:index]
	}
	if len(p.Content) > 150 {
		return p.Content[0:150]
	}
	return p.Content
}

func (p Post) Permalink() string {
	return fmt.Sprintf("/%d/%d/%s", p.Created.Year(), p.Created.Month(), p.Slug)
}
