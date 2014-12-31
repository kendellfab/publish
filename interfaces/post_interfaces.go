package interfaces

import (
	"database/sql"
	"fmt"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

type DbPostRepo struct {
	db         *sql.DB
	authorRepo domain.UserRepo
	catRepo    domain.CategoryRepo
	cache      *PostCache
}

func NewDbPostRepo(db *sql.DB, ar domain.UserRepo, cr domain.CategoryRepo) domain.PostRepo {
	postRepo := &DbPostRepo{db: db, authorRepo: ar, catRepo: cr}
	cache, err := NewPostCache(100)
	if err != nil {
		log.Fatal(err)
	}
	postRepo.cache = cache
	postRepo.init()
	return postRepo
}

func (repo *DbPostRepo) init() {
	if _, err := repo.db.Exec(CREATE_POST); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
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
	res, err := repo.db.Exec("INSERT INTO post(title, slug, author, created, content, type, published, tags, category, day, month, year, series) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", post.Title, post.Slug, authorId, createdStr, post.Content, post.ContentType, published, tagsStr, 1, day, month, year, post.SeriesId)
	if err == nil {
		if id, idErr := res.LastInsertId(); idErr == nil {
			post.Id = id
		}
	}
	return err
}

func (repo *DbPostRepo) Update(post *domain.Post) error {
	tags, _ := domain.SerializeTags(post.Tags)
	sql := "UPDATE post SET title=?, slug=?, content=?, type=?, tags=?, category=?, series=? WHERE id = ?"
	catId := 1
	if post.Category != nil {
		catId = post.Category.Id
	}
	_, err := repo.db.Exec(sql, post.Title, post.Slug, post.Content, post.ContentType, tags, catId, post.SeriesId, post.Id)

	return err
}

func (repo *DbPostRepo) Publish(id int64) error {
	sql := "UPDATE post SET published = 1 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPostRepo) UnPublish(id int64) error {
	sql := "UPDATE post SET published = 0 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPostRepo) FindById(id int64) (*domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post WHERE id=?"
	row := repo.db.QueryRow(sql, id)
	return repo.scanPost(row)
}

func (repo *DbPostRepo) FindByIdString(id string) (*domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post WHERE id=?"
	row := repo.db.QueryRow(sql, id)
	return repo.scanPost(row)
}

func (repo *DbPostRepo) FindBySlug(slug string) (*domain.Post, error) {
	if post, ok := repo.cache.Get(slug); ok {
		return post, nil
	}
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post WHERE slug=?"
	row := repo.db.QueryRow(sql, slug)
	return repo.scanPost(row)
}

func (repo *DbPostRepo) FindByCategory(category *domain.Category, offset, limit int) ([]*domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post WHERE category=? AND published = 1 ORDER BY created DESC LIMIT ? OFFSET ?;"
	rows, qError := repo.db.Query(sql, category.Id, limit, offset)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPosts(rows)
	return posts, nil
}

func (repo *DbPostRepo) FindAll() ([]*domain.Post, error) {
	sql := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post ORDER BY created DESC;"
	rows, qError := repo.db.Query(sql)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPosts(rows)
	return posts, nil
}

func (repo *DbPostRepo) getPublishedIds(offset, limit int) ([]string, error) {
	sel := "SELECT slug FROM post WHERE published = 1 ORDER BY created DESC LIMIT ? OFFSET ?;"
	rows, qErr := repo.db.Query(sel, limit, offset)
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

func (repo *DbPostRepo) FindPublished(offset, limit int) ([]*domain.Post, error) {
	posts := make([]*domain.Post, 0)

	ids, err := repo.getPublishedIds(offset, limit)
	faults := make([]interface{}, 0)
	if err == nil {
		for _, id := range ids {
			if post, ok := repo.cache.Get(id); ok {
				posts = append(posts, post)
			} else {
				faults = append(faults, id)
			}
		}
	}

	if len(faults) == 0 {
		return posts, nil
	}

	var rows *sql.Rows
	var qError error

	if len(faults) == len(ids) {
		sel := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post WHERE published = 1 ORDER BY created DESC LIMIT ? OFFSET ?;"
		rows, qError = repo.db.Query(sel, limit, offset)
	} else {
		sel := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post WHERE slug IN(%s) ORDER BY created DESC;"
		phs := fmt.Sprintf(sel, GetPlaceholders(faults))
		rows, qError = repo.db.Query(phs, faults...)
	}

	if qError != nil {
		return nil, qError
	}

	results := repo.scanPosts(rows)
	return append(posts, results...), nil
}

func (repo *DbPostRepo) GetForSeries(seriesId string) ([]*domain.Post, error) {
	sel := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post WHERE published = 1 AND series = ? ORDER BY created DESC;"
	rows, qErr := repo.db.Query(sel, seriesId)
	if qErr != nil {
		return nil, qErr
	}
	return repo.scanPosts(rows), nil
}

func (repo *DbPostRepo) FindByYearMonth(year, month string) ([]*domain.Post, error) {
	sel := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post WHERE published = 1 AND year = ? AND month = ? ORDER BY created DESC;"
	rows, qError := repo.db.Query(sel, year, month)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPosts(rows)
	return posts, nil
}

func (repo *DbPostRepo) FindDashboard(offset, limit int) ([]*domain.Post, error) {
	sel := "SELECT id, title, slug, author, created, content, type, published, tags, category, series FROM post ORDER BY created DESC LIMIT ? OFFSET ?;"
	rows, qError := repo.db.Query(sel, limit, offset)
	if qError != nil {
		return nil, qError
	}
	posts := repo.scanPosts(rows)
	return posts, nil
}

func (repo *DbPostRepo) Delete(id string) error {
	sql := "DELETE FROM post WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func (repo *DbPostRepo) PublishedCount() (int, error) {
	sel := "SELECT count(*) FROM post WHERE published = 1;"
	row := repo.db.QueryRow(sel)
	var count int
	scanErr := row.Scan(&count)
	if scanErr != nil {
		return 0, scanErr
	}
	return count, nil
}

func (repo *DbPostRepo) PublishedCountCategory(catId int) (int, error) {
	sel := "SELECT count(*) FROM post WHERE published = 1 AND category = ?;"
	row := repo.db.QueryRow(sel, catId)
	var count int
	scanErr := row.Scan(&count)
	if scanErr != nil {
		return 0, scanErr
	}
	return count, nil
}

func (repo *DbPostRepo) scanPost(row *sql.Row) (*domain.Post, error) {
	var post domain.Post
	var authorId int64
	var createString string
	var tagString string
	var published int
	var categoryId int
	scanErr := row.Scan(&post.Id, &post.Title, &post.Slug, &authorId, &createString, &post.Content, &post.ContentType, &published, &tagString, &categoryId, &post.SeriesId)
	if scanErr != nil {
		return nil, scanErr
	}

	if auth, aErr := repo.authorRepo.FindByIdInt(authorId); aErr == nil {
		post.Author = *auth
	}

	if cat, cErr := repo.catRepo.FindById(categoryId); cErr == nil {
		post.Category = cat
	}

	post.Created, _ = time.Parse(time.RFC3339, createString)
	post.Published = false
	if published == 1 {
		post.Published = true
	}
	repo.cache.Add(post.Slug, &post)
	return &post, nil

}

func (repo *DbPostRepo) scanPosts(rows *sql.Rows) []*domain.Post {
	posts := make([]*domain.Post, 0)
	rows.Next()
	for {
		var post domain.Post
		var authorId int64
		var createString string
		var tagsString string
		var published int
		var categoryId int
		scanErr := rows.Scan(&post.Id, &post.Title, &post.Slug, &authorId, &createString, &post.Content, &post.ContentType, &published, &tagsString, &categoryId, &post.SeriesId)

		if scanErr == nil {
			author, err := repo.authorRepo.FindByIdInt(authorId)

			if err == nil {
				post.Author = *author
			}

			cat, err := repo.catRepo.FindById(categoryId)
			if err == nil {
				post.Category = cat
			} else {
				post.Category = nil
			}

			post.Created, _ = time.Parse(time.RFC3339, createString)
			post.Published = false
			if published == 1 {
				post.Published = true
			}
			posts = append(posts, &post)
			repo.cache.Add(post.Slug, &post)
		} else {
			log.Println("Scan Posts:", scanErr.Error())
		}

		if !rows.Next() {
			break
		}
	}
	return posts
}
