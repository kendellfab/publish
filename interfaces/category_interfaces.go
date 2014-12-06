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
	exec := `CREATE TABLE "category" (
"id" INTEGER PRIMARY KEY NOT NULL,
"title" TEXT NOT NULL,
"slug" TEXT NOT NULL,
"created" TEXT
)`
	if _, err := repo.db.Exec(exec); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
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

func (repo *DbCategoryRepo) GetAll() (*[]domain.Category, error) {
	rows, rowError := repo.db.Query("SELECT id, title, slug, created FROM category")
	if rowError != nil {
		return nil, rowError
	}
	cats := scanCategory(rows)
	return &cats, nil
}

func (repo *DbCategoryRepo) DeleteById(id int) error {
	sql := "DELETE FROM category WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbCategoryRepo) GetCategoryPostCount() (*[]domain.CategoryCount, error) {
	sql := `select c.title, c.slug, count(*)
from post p
join category c on p.category = c.id
group by category;`
	row, qError := repo.db.Query(sql)
	if qError != nil {
		return nil, qError
	}
	counts := make([]domain.CategoryCount, 0)
	for {
		var count domain.CategoryCount
		scanErr := row.Scan(&count.Title, &count.Slug, &count.Count)
		if scanErr == nil {
			counts = append(counts, count)
		}
		if !row.Next() {
			break
		}
	}
	return &counts, nil
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
