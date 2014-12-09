package interfaces

import (
	"database/sql"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

type DbCategoryRepo struct {
	db *sql.DB
}

func NewDbCategoryRepo(db *sql.DB) domain.CategoryRepo {
	catRepo := &DbCategoryRepo{db: db}
	catRepo.init()
	return catRepo
}

func (repo *DbCategoryRepo) init() {
	if _, err := repo.db.Exec(CREATE_CATEGORY); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
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
	var category domain.Category
	var dateStr string
	row := repo.db.QueryRow("SELECT id, title, slug, created FROM category WHERE id=?", id)
	scanErr := row.Scan(&category.Id, &category.Title, &category.Slug, &dateStr)
	if scanErr == nil {
		// switch {
		// case scanErr == sql.ErrNoRows:
		// 	return nil, scanErr
		// case scanErr != nil:
		// 	repo.logger.LogError(scanErr)
		// 	return nil, scanErr
		// }
		date, _ := time.Parse(time.RFC3339, dateStr)
		category.Created = date
		return &category, nil
	} else {
		return nil, scanErr
	}
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

func (repo *DbCategoryRepo) GetAll() ([]*domain.Category, error) {
	rows, qError := repo.db.Query("SELECT c.id, c.title, c.slug, c.created, count(*) FROM post p join category c on p.category = c.id where p.published = 1 group by category;")
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

func scanCategory(rows *sql.Rows) []domain.Category {
	cats := make([]domain.Category, 0)
	for {
		var category domain.Category
		var dateStr string
		scanErr := rows.Scan(&category.Id, &category.Title, &category.Slug, &dateStr)
		if scanErr == nil {
			date, _ := time.Parse(time.RFC3339, dateStr)
			category.Created = date
			cats = append(cats, category)
		}
		if !rows.Next() {
			break
		}
	}

	return cats
}
