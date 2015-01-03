package interfaces

import (
	"database/sql"
	"fmt"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

type DbPageRepo struct {
	db    *sql.DB
	cache *PageCache
}

func NewDbPageRepo(db *sql.DB) domain.PageRepo {
	pageRepo := &DbPageRepo{db: db}
	cache, err := NewPageCache(100)
	if err != nil {
		log.Fatal(err)
	}
	pageRepo.cache = cache
	pageRepo.init()
	return pageRepo
}

func (repo *DbPageRepo) init() {
	if _, err := repo.db.Exec(CREATE_PAGE); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbPageRepo) Store(page *domain.Page) error {
	createdString := domain.SerializeDate(page.Created)
	res, err := repo.db.Exec("INSERT INTO page(title, slug, created, content, published) VALUES(?, ?, ?, ?, 0)", page.Title, page.Slug, createdString, page.Content)
	if err != nil {
		return err
	}
	if id, idErr := res.LastInsertId(); idErr == nil {
		page.Id = id
	}
	return nil
}

func (repo *DbPageRepo) Update(page *domain.Page) error {
	createdString := domain.SerializeDate(page.Created)
	published := 0
	if page.Published {
		published = 1
	}
	_, err := repo.db.Exec("UPDATE page SET title = ?, slug = ?, created = ?, content = ?, published = ? WHERE id = ?", page.Title, page.Slug, createdString, page.Content, published, page.Id)
	return err
}

func (repo *DbPageRepo) FindById(id string) (*domain.Page, error) {
	raw := "SELECT id, title, slug, created, content, published FROM page WHERE id = ?"
	row := repo.db.QueryRow(raw, id)
	return repo.parseRow(row)
}

func (repo *DbPageRepo) FindBySlug(slug string) (*domain.Page, error) {
	if page, ok := repo.cache.Get(slug); ok {
		return page, nil
	}
	raw := "SELECT id, title, slug, created, content, published FROM page WHERE slug = ?"
	row := repo.db.QueryRow(raw, slug)
	return repo.parseRow(row)
}

func (repo *DbPageRepo) FindAll() ([]*domain.Page, error) {
	sql := "SELECT id, title, slug, created, content, published FROM page"
	rows, qError := repo.db.Query(sql)
	if qError != nil {
		return nil, qError
	}
	pages := repo.parseRows(rows)
	return pages, nil
}

func (repo *DbPageRepo) getPublishedIds() ([]string, error) {
	sel := "SELECT slug FROM page WHERE published = 1;"
	rows, qErr := repo.db.Query(sel)
	if qErr != nil {
		return nil, qErr
	}
	slugs := make([]string, 0)
	rows.Next()
	for {
		var slug string
		sErr := rows.Scan(&slug)
		if sErr == nil {
			slugs = append(slugs, slug)
		}
		if !rows.Next() {
			break
		}
	}
	return slugs, nil
}

func (repo *DbPageRepo) FindAllPublished() ([]*domain.Page, error) {
	pages := make([]*domain.Page, 0)

	slugs, err := repo.getPublishedIds()
	faults := make([]interface{}, 0)
	if err == nil {
		for _, slug := range slugs {
			if page, ok := repo.cache.Get(slug); ok {
				pages = append(pages, page)
			} else {
				faults = append(faults, slug)
			}
		}
	}

	if len(faults) == 0 {
		return pages, nil
	}

	var rows *sql.Rows
	var qError error

	if len(faults) == len(slugs) {
		sel := "SELECT id, title, slug, created, content, published FROM page WHERE published = 1"
		rows, qError = repo.db.Query(sel)
	} else {
		sel := "SELECT id, title, slug, created, content, published FROM page WHERE slug IN(%s);"
		phs := fmt.Sprintf(sel, GetPlaceholders(faults))
		rows, qError = repo.db.Query(phs, faults...)
	}

	if qError != nil {
		return nil, qError
	}
	results := repo.parseRows(rows)
	return append(pages, results...), nil
}

func (repo *DbPageRepo) Publish(id string) error {
	sql := "UPDATE page SET published = 1 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPageRepo) UnPublish(id string) error {
	sql := "UPDATE page SET published = 0 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPageRepo) Delete(id string) error {
	sql := "DELETE FROM page WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPageRepo) parseRows(rows *sql.Rows) []*domain.Page {
	pages := make([]*domain.Page, 0)
	for {
		var page domain.Page
		var createdStr string
		var published int
		scanErr := rows.Scan(&page.Id, &page.Title, &page.Slug, &createdStr, &page.Content, &published)

		if scanErr == nil {
			date, _ := time.Parse(time.RFC3339, createdStr)
			page.Created = date
			page.Published = false
			if published == 1 {
				page.Published = true
			}
			repo.cache.Add(page.Slug, &page)
			pages = append(pages, &page)
		}
		if !rows.Next() {
			break
		}
	}

	return pages
}

func (repo *DbPageRepo) parseRow(row *sql.Row) (*domain.Page, error) {
	var page domain.Page
	var createdStr string
	var published int
	scanErr := row.Scan(&page.Id, &page.Title, &page.Slug, &createdStr, &page.Content, &published)

	switch {
	case scanErr == sql.ErrNoRows:
		return nil, scanErr
	case scanErr != nil:
		return nil, scanErr
	}

	date, _ := time.Parse(time.RFC3339, createdStr)
	page.Created = date
	page.Published = false
	if published == 1 {
		page.Published = true
	}

	return &page, nil
}
