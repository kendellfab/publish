package interfaces

import (
	"database/sql"
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

func (repo *DbSeriesRepo) GetAll() ([]*domain.Series, error) {
	sel := "SELECT id, title, slug, created, description FROM series;"
	rows, qErr := repo.db.Query(sel)
	if qErr != nil {
		return nil, qErr
	}

	series := make([]*domain.Series, 0)
	rows.Next()
	for {
		var s domain.Series
		var created string
		sErr := rows.Scan(&s.Id, &s.Title, &s.Slug, &created, &s.Description)
		if sErr == nil {
			s.Created, _ = time.Parse(time.RFC3339, created)
			series = append(series, &s)
		}
		if !rows.Next() {
			break
		}
	}

	return series, nil
}

func (repo *DbSeriesRepo) GetSeries(id string) (*domain.Series, error) {
	sel := "SELECT id, title, slug, created, description FROM series WHERE id = ?;"
	row := repo.db.QueryRow(sel, id)

	var s domain.Series
	var created string
	sErr := row.Scan(&s.Id, &s.Title, &s.Slug, &created, &s.Description)
	if sErr != nil {
		return nil, sErr
	}

	s.Created, _ = time.Parse(time.RFC3339, created)

	return &s, nil
}