package interfaces

import (
	"database/sql"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

type DbPostRepo struct {
	db         *sql.DB
	authorRepo domain.UserRepo
	catRepo    domain.CategoryRepo
}

func NewDbPostRepo(db *sql.DB, ar domain.UserRepo, cr domain.CategoryRepo) domain.PostRepo {
	postRepo := &DbPostRepo{db: db, authorRepo: ar, catRepo: cr}
	postRepo.init()
	return postRepo
}

func (repo *DbPostRepo) init() {
	exec := `CREATE TABLE post (
"id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
"title" TEXT NOT NULL,
"slug" TEXT NOT NULL,
"author" INTEGER NOT NULL,
"created" TEXT NOT NULL,
"content" TEXT NOT NULL,
"type" TEXT,
"published" INTEGER,
"tags" TEXT,
"category" INTEGER,
"day" INTEGER,
"month" INTEGER,
"year" INTEGER
)`

	if _, err := repo.db.Exec(exec); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbPostRepo) Store(post *domain.Post) error {
	day, month, year := domain.DateComponents(post.Created)
	published := 0
	if post.Published {
		published = 1
	}
	createdStr := domain.SerializeDate(post.Created)
	tagsStr, _ := domain.SerializeTags(post.Tags)
	authorId := post.AuthorId
	_, err := repo.db.Exec("INSERT INTO post(title, slug, author, created, content, type, published, tags, category, day, month, year) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", post.Title, post.Slug, authorId, createdStr, post.Content, post.ContentType, published, tagsStr, post.Category.Id, day, month, year)
	return err
}

func (repo *DbPostRepo) Update(post *domain.Post) error {
	tags, _ := domain.SerializeTags(post.Tags)
	sql := "UPDATE post SET title=?, slug=?, content=?, type=?, tags=?, category=? WHERE id = ?"
	res, err := repo.db.Exec(sql, post.Title, post.Slug, post.Content, post.ContentType, tags, post.Category.Id, post.Id)

	if err == nil {
		if id, idErr := res.LastInsertId(); idErr == nil {
			post.Id = id
		}
	}

	return err
}

func (repo *DbPostRepo) Publish(id int) error {
	sql := "UPDATE post SET published = 1 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPostRepo) UnPublish(id int) error {
	sql := "UPDATE post SET published = 0 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPostRepo) FindById(id int64) (*domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category FROM post WHERE id=?"
	rows, qError := repo.db.Query(sql, id)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPost(rows)
	return &posts[0], nil
}

func (repo *DbPostRepo) FindByIdString(id string) (*domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category FROM post WHERE id=?"
	rows, qError := repo.db.Query(sql, id)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPost(rows)
	return &posts[0], nil
}

func (repo *DbPostRepo) FindBySlug(slug string) (*domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category FROM post WHERE slug=?"
	rows, qError := repo.db.Query(sql, slug)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPost(rows)
	return &posts[0], nil
}

func (repo *DbPostRepo) FindByCategory(category *domain.Category) (*[]domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category FROM post WHERE category=?"
	rows, qError := repo.db.Query(sql, category.Id)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPost(rows)
	return &posts, nil
}

func (repo *DbPostRepo) FindAll() (*[]domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category FROM post"
	rows, qError := repo.db.Query(sql)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPost(rows)
	return &posts, nil
}

func (repo *DbPostRepo) Delete(id int) error {
	sql := "DELETE FROM post WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPostRepo) scanPost(rows *sql.Rows) []domain.Post {
	posts := make([]domain.Post, 0)
	authors := make(map[int64]domain.User)
	cats := make(map[int]domain.Category)
	for {
		var post domain.Post
		var authorId int64
		var createString string
		var tagsString string
		var published int
		var categoryId int
		scanErr := rows.Scan(&post.Id, &post.Title, &post.Slug, &authorId, &createString, &post.Content, &post.ContentType, &published, &tagsString, &categoryId)

		if scanErr == nil {
			if a, ok := authors[authorId]; ok {
				post.Author = a
			} else {
				author, err := repo.authorRepo.FindByIdInt(authorId)

				if err == nil {
					post.Author = *author
					authors[authorId] = *author
				}
			}

			if c, ok := cats[categoryId]; ok {
				post.Category = c
			} else {
				cat, err := repo.catRepo.FindById(categoryId)
				if err == nil {
					post.Category = *cat
					cats[categoryId] = *cat
				}
			}

			post.Created, _ = time.Parse(time.RFC3339, createString)
			post.Published = false
			if published == 1 {
				post.Published = true
			}
			posts = append(posts, post)
		}

		if !rows.Next() {
			break
		}
	}
	return posts
}
