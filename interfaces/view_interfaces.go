package interfaces

import (
	"database/sql"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
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
	ins := "INSERT INTO view(who, at, type, target) VALUES(?, ?, ?, ?);"
	at := v.At.Format(time.RFC3339)
	_, err := repo.db.Exec(ins, v.Who, at, int(v.TargetType), v.Target)
	return err
}

func (repo *DbViewRepo) GetType(t domain.TargetType) ([]*domain.View, error) {
	return nil, nil
}

func (repo *DbViewRepo) GetTypeTarget(t domain.TargetType, target string) ([]*domain.View, error) {
	sel := "SELECT id, who, at, type, target FROM view WHERE type = ? AND target = ? ORDER BY at DESC;"
	rows, qErr := repo.db.Query(sel, t, target)
	if qErr != nil {
		return nil, qErr
	}
	return repo.scanRows(rows)
}

func (repo *DbViewRepo) scanRows(rows *sql.Rows) ([]*domain.View, error) {
	views := make([]*domain.View, 0)
	for {
		var view domain.View
		var at string
		sErr := rows.Scan(&view.Id, &view.Who, &at, &view.TargetType, &view.Target)
		if sErr == nil {
			view.At, _ = time.Parse(time.RFC3339, at)
			views = append(views, &view)
		}
		if !rows.Next() {
			break
		}
	}
	return views, nil
}
