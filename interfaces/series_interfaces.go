package interfaces

import (
	"database/sql"
	"fmt"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

type DbSeriesRepo struct {
	db       *sql.DB
	postRepo domain.PostRepo
}

func NewDbSeriesRepo(db *sql.DB, pr domain.PostRepo) domain.SeriesRepo {
	repo := &DbSeriesRepo{db: db, postRepo: pr}
	repo.init()
	return repo
}

func (repo *DbSeriesRepo) init() {
	if _, err := repo.db.Exec(CREATE_SERIES); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbSeriesRepo) Store(s *domain.Series) error {
	ins := "INSERT INTO series(title, slug, created, description) VALUES(?, ?, ?, ?);"
	res, err := repo.db.Exec(ins, s.Title, s.Slug, s.Created.Format(time.RFC3339), s.Description)
	if id, idErr := res.LastInsertId(); idErr == nil {
		s.Id = id
	}
	return err
}

func (repo *DbSeriesRepo) Count() (int, error) {
	sel := "SELECT count(*) FROM series;"
	row := repo.db.QueryRow(sel)
	var count int
	sErr := row.Scan(&count)
	if sErr != nil {
		return 0, sErr
	}
	return count, nil
}

func (repo *DbSeriesRepo) Update(s *domain.Series) error {
	up := "UPDATE series SET title=?, slug=?, created=?, description=? WHERE id = ?;"
	_, err := repo.db.Exec(up, s.Title, s.Slug, s.Created.Format(time.RFC3339), s.Description, s.Id)
	return err
}

func (repo *DbSeriesRepo) GetAll() ([]*domain.Series, error) {
	sel := "SELECT id, title, slug, created, description FROM series ORDER BY created DESC;"
	rows, qErr := repo.db.Query(sel)
	if qErr != nil {
		return nil, qErr
	}

	return repo.scanSeries(rows), nil
}

func (repo *DbSeriesRepo) GetSeriesLimit(offset, limit int) ([]*domain.Series, error) {
	sel := "SELECT id, title, slug, created, description FROM series ORDER BY created DESC LIMIT ? OFFSET ?;"
	rows, qErr := repo.db.Query(sel, limit, offset)
	if qErr != nil {
		return nil, qErr
	}
	return repo.scanSeries(rows), nil
}

func (repo *DbSeriesRepo) GetSeries(id string) (*domain.Series, error) {
	sel := "SELECT id, title, slug, created, description FROM series WHERE id = ?;"
	row := repo.db.QueryRow(sel, id)

	return repo.scanSingleSeries(row)
}

func (repo *DbSeriesRepo) GetSeriesWithSlug(slug string) (*domain.Series, error) {
	sel := "SELECT id, title, slug, created, description FROM series WHERE slug = ?;"
	row := repo.db.QueryRow(sel, slug)

	return repo.scanSingleSeries(row)
}

func (repo *DbSeriesRepo) scanSingleSeries(row *sql.Row) (*domain.Series, error) {
	var s domain.Series
	var created string
	sErr := row.Scan(&s.Id, &s.Title, &s.Slug, &created, &s.Description)
	if sErr != nil {
		return nil, sErr
	}

	s.Created, _ = time.Parse(time.RFC3339, created)

	if posts, pErr := repo.postRepo.GetForSeries(fmt.Sprintf("%d", s.Id)); pErr == nil {
		s.Posts = posts
	} else {
		log.Println("Series:", pErr)
	}

	return &s, nil
}

func (repo *DbSeriesRepo) scanSeries(rows *sql.Rows) []*domain.Series {
	series := make([]*domain.Series, 0)
	rows.Next()
	for {
		var s domain.Series
		var created string
		sErr := rows.Scan(&s.Id, &s.Title, &s.Slug, &created, &s.Description)
		if sErr == nil {
			s.Created, _ = time.Parse(time.RFC3339, created)
			if posts, pErr := repo.postRepo.GetForSeries(fmt.Sprintf("%d", s.Id)); pErr == nil {
				s.Posts = posts
			}
			series = append(series, &s)
		}
		if !rows.Next() {
			break
		}
	}
	return series
}
