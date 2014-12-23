package interfaces

import (
	"database/sql"
	"fmt"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

var uncategorized = &domain.Category{Title: "Uncategorized", Slug: "uncategorized"}

type DbCategoryRepo struct {
	db    *sql.DB
	cache *CategoryCache
}

func NewDbCategoryRepo(db *sql.DB) domain.CategoryRepo {
	catRepo := &DbCategoryRepo{db: db}
	cache, err := NewCategoryCache(25)
	if err != nil {
		log.Fatal(err)
	}
	catRepo.cache = cache
	catRepo.init()
	return catRepo
}

func (repo *DbCategoryRepo) init() {
	if _, err := repo.db.Exec(CREATE_CATEGORY); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}

	countQ := "SELECT count(*) FROM category;"
	row := repo.db.QueryRow(countQ)
	var count int
	scanErr := row.Scan(&count)
	if (scanErr == nil || scanErr == sql.ErrNoRows) && count == 0 {
		log.Println("Setting up default category.")
		uncategorized.Created = time.Now()
		repo.Store(uncategorized)
	} else if scanErr != nil {
		log.Fatal("Error", scanErr, "Count", count)
	}
}

func (repo *DbCategoryRepo) Store(category *domain.Category) error {
	dateStr := category.Created.Format(time.RFC3339)
	if category.Slug == "" {
		category.GenerateSlug()
	}
	sql := "INSERT INTO category(title, slug, created) VALUES(?, ?, ?)"
	_, err := repo.db.Exec(sql, category.Title, category.Slug, dateStr)
	return err
}

func (repo *DbCategoryRepo) FindById(id int) (*domain.Category, error) {
	if cat, ok := repo.cache.Get(fmt.Sprintf("%d", id)); ok {
		return cat, nil
	}
	var category domain.Category
	var dateStr string
	row := repo.db.QueryRow("SELECT id, title, slug, created FROM category WHERE id=?", id)
	scanErr := row.Scan(&category.Id, &category.Title, &category.Slug, &dateStr)
	if scanErr != nil {
		return nil, scanErr
	}

	date, _ := time.Parse(time.RFC3339, dateStr)
	category.Created = date

	repo.cache.Add(fmt.Sprintf("%d", id), &category)

	return &category, nil
}

func (repo *DbCategoryRepo) FindByTitle(title string) (*domain.Category, error) {
	var category domain.Category
	var dateStr string
	row := repo.db.QueryRow("SELECT id, title, slug, created FROM category WHERE title=?", title)
	scanErr := row.Scan(&category.Id, &category.Title, &category.Slug, &dateStr)
	if scanErr == nil {
		date, _ := time.Parse(time.RFC3339, dateStr)
		category.Created = date
		return &category, nil
	} else {
		return nil, scanErr
	}
}

func (repo *DbCategoryRepo) FindBySlug(slug string) (*domain.Category, error) {
	var category domain.Category
	var dateStr string
	row := repo.db.QueryRow("SELECT id, title, slug, created FROM category WHERE slug=?", slug)
	scanErr := row.Scan(&category.Id, &category.Title, &category.Slug, &dateStr)
	if scanErr == nil {
		date, _ := time.Parse(time.RFC3339, dateStr)
		category.Created = date
		return &category, nil
	} else {
		return nil, scanErr
	}
}

func (repo *DbCategoryRepo) getCatIds() ([]string, error) {
	sel := "SELECT id FROM category;"
	rows, qErr := repo.db.Query(sel)
	if qErr != nil {
		return nil, qErr
	}
	ids := make([]string, 0)
	rows.Next()
	for {
		var id string
		sErr := rows.Scan(&id)
		if sErr == nil {
			ids = append(ids, id)
		}
		if !rows.Next() {
			break
		}
	}
	return ids, nil
}

func (repo *DbCategoryRepo) GetAll() ([]*domain.Category, error) {
	cats := make([]*domain.Category, 0)
	ids, err := repo.getCatIds()
	faults := make([]interface{}, 0)
	if err == nil {
		for _, id := range ids {
			if cat, ok := repo.cache.Get(id); ok {
				cats = append(cats, cat)
			} else {
				faults = append(faults, id)
			}
		}
	}

	if len(faults) == 0 {
		return cats, nil
	}

	var rows *sql.Rows
	var qError error
	if len(faults) == len(ids) {
		rows, qError = repo.db.Query("SELECT id, title, slug, created FROM category")
	} else {
		sel := "SELECT id, title, slug, created FROM category WHERE id IN(%s);"
		phs := fmt.Sprintf(sel, GetPlaceholders(faults))
		rows, qError = repo.db.Query(phs, faults...)
	}

	if qError != nil {
		return nil, qError
	}
	results := repo.scanCategory(rows)
	return append(cats, results...), nil
}

func (repo *DbCategoryRepo) GetAllCount() ([]*domain.Category, error) {
	rows, qError := repo.db.Query("SELECT c.id, c.title, c.slug, c.created, count(*) FROM category c join post p on p.category = c.id where p.published = 1 group by category;")
	if qError != nil {
		return nil, qError
	}
	cats := make([]*domain.Category, 0)
	for {
		var cat domain.Category
		var created string
		scanErr := rows.Scan(&cat.Id, &cat.Title, &cat.Slug, &created, &cat.Count)
		if scanErr == nil {
			cat.Created, _ = time.Parse(time.RFC3339, created)
			cats = append(cats, &cat)
		}
		if !rows.Next() {
			break
		}
	}
	return cats, nil
}

func (repo *DbCategoryRepo) DeleteById(id int) error {
	sql := "DELETE FROM category WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbCategoryRepo) scanCategory(rows *sql.Rows) []*domain.Category {
	cats := make([]*domain.Category, 0)
	for {
		var category domain.Category
		var dateStr string
		scanErr := rows.Scan(&category.Id, &category.Title, &category.Slug, &dateStr)
		if scanErr == nil {
			date, _ := time.Parse(time.RFC3339, dateStr)
			category.Created = date
			cats = append(cats, &category)
		}
		if !rows.Next() {
			break
		}
	}

	return cats
}
