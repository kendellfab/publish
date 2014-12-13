// +build sqlite

package infrastructure

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"errors"
	"github.com/kendellfab/publish/domain"
)

func init() {
	CurrentDb = DbSqlite
}

func ConnectDb(config *domain.Config) (*sql.DB, error) {
	if config == nil || config.Sqlite == nil {
		return nil, errors.New("Publish: Database config required!")
	}
	return sql.Open("sqlite3", config.Sqlite.Path)
}
