package domain

import (
	"strings"
	"time"
)

type TimeSeries struct {
	When  time.Time
	Posts []*Post
}

type SeriesRepo interface {
	Store(s *Series) error
	GetAll() ([]*Series, error)
	GetSeries(id string) (*Series, error)
}

type Series struct {
	Id          int64
	Title       string
	Slug        string
	Created     time.Time
	Description string
	Posts       []*Post
}

func (s *Series) GenerateSlug() {
	if s.Slug == "" {
		slug := s.Title
		slug = strings.Replace(slug, " ", SLUG_SPACER, -1)
		slug = strings.ToLower(slug)
		s.Slug = slug
	}
}
