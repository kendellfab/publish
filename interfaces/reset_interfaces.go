package interfaces

import (
	"database/sql"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

type DbResetRepo struct {
	db *sql.DB
}

func NewDbResetRepo(db *sql.DB) domain.ResetRepo {
	repo := &DbResetRepo{db: db}
	repo.init()
	return repo
}

func (repo *DbResetRepo) init() {
	if _, err := repo.db.Exec(CREATE_RESET); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbResetRepo) Store(r *domain.Reset) error {
	ins := "INSERT INTO reset(who, created, expires, token) VALUES(?, ?, ?, ?);"
	_, err := repo.db.Exec(ins, r.UserId, r.Created.Format(time.RFC3339), r.Expires.Format(time.RFC3339), r.Token)
	return err
}

func (repo *DbResetRepo) FindByToken(token string) (*domain.Reset, error) {
	sel := "SELECT id, who, created, expires, token FROM reset WHERE token = ?;"
	row := repo.db.QueryRow(sel, token)
	var reset domain.Reset
	var created string
	var expires string
	sErr := row.Scan(&reset.Id, &reset.UserId, &created, &expires, &reset.Token)
	if sErr != nil {
		return nil, sErr
	}
	reset.Created, _ = time.Parse(time.RFC3339, created)
	reset.Expires, _ = time.Parse(time.RFC3339, expires)
	return &reset, nil
}

func (repo *DbResetRepo) CleanExpired() error {
	return nil
}
