package interfaces

import (
	"database/sql"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

type DbPageRepo struct {
	db *sql.DB
}

func NewDbPageRepo(db *sql.DB) domain.PageRepo {
	pageRepo := &DbPageRepo{db: db}
	pageRepo.init()
	return pageRepo
}

func (repo *DbPageRepo) init() {
	exec := `CREATE TABLE "page" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
"title" TEXT NOT NULL,
"slug" TEXT NOT NULL,
"created" TEXT NOT NULL,
"content" TEXT NOT NULL,
"published" INTEGER
)`
	if _, err := repo.db.Exec(exec); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbPageRepo) Store(page *domain.Page) error {
	createdString := domain.SerializeDate(page.Created)
	_, err := repo.db.Exec("INSERT INTO page(title, slug, created, content, published) VALUES(?, ?, ?, ?, 0)", page.Title, page.Slug, createdString, page.Content)
	return err
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

func (repo *DbPageRepo) FindById(id int) (*domain.Page, error) {
	raw := "SELECT id, title, slug, created, content, published FROM page WHERE id = ?"
	row := repo.db.QueryRow(raw, id)
	return repo.parseRow(row)
}

func (repo *DbPageRepo) FindBySlug(slug string) (*domain.Page, error) {
	raw := "SELECT id, title, slug, created, content, published FROM page WHERE slug = ?"
	row := repo.db.QueryRow(raw, slug)
	return repo.parseRow(row)
}

func (repo *DbPageRepo) FindAll() (*[]domain.Page, error) {
	sql := "SELECT id, title, slug, created, content, published FROM page"
	rows, qError := repo.db.Query(sql)
	if qError != nil {
		return nil, qError
	}
	pages := repo.parseRows(rows)
	return &pages, nil
}

func (repo *DbPageRepo) FindAllPublished() (*[]domain.Page, error) {
	sql := "SELECT id, title, slug, created, content, published FROM page WHERE published = 1"
	rows, qError := repo.db.Query(sql)
	if qError != nil {
		return nil, qError
	}
	pages := repo.parseRows(rows)
	return &pages, nil
}

func (repo *DbPageRepo) Publish(id int) error {
	sql := "UPDATE page SET published = 1 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPageRepo) UnPublish(id int) error {
	sql := "UPDATE page SET published = 0 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPageRepo) Delete(id int) error {
	sql := "DELETE FROM page WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPageRepo) parseRows(rows *sql.Rows) []domain.Page {
	pages := make([]domain.Page, 0)
	for {
		var page domain.Page
		var createdStr string
		var published int
		scanErr := rows.Scan(&page.Id, &page.Title, &page.Slug, &createdStr, &page.Content, &published)

		if scanErr == nil {
			date, _ := time.Parse(domain.DATE_STORAGE_FORMAT, createdStr)
			page.Created = date
			page.Published = false
			if published == 1 {
				page.Published = true
			}
			pages = append(pages, page)
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

	date, _ := time.Parse(domain.DATE_STORAGE_FORMAT, createdStr)
	page.Created = date
	page.Published = false
	if published == 1 {
		page.Published = true
	}

	return &page, nil
}
