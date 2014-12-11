package interfaces

import (
	"database/sql"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
)

type DbViewRepo struct {
	db *sql.DB
}

func NewDbViewRepo(db *sql.DB) domain.ViewRepo {
	repo := &DbViewRepo{db: db}
	repo.init()
	return repo
}

func (repo *DbViewRepo) init() {
	if _, err := repo.db.Exec(CREATE_VIEW); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbViewRepo) Store(v *domain.View) error {
	return nil
}

func (repo *DbViewRepo) GetType(t domain.TargetType) ([]*domain.View, error) {
	return nil, nil
}

func (repo *DbViewRepo) GetTypeTarget(t domain.TargetType, target string) ([]*domain.View, error) {
	return nil, nil
}
