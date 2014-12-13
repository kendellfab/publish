// +build sqlite

package infrastructure

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"errors"
	"github.com/kendellfab/publish/domain"
	"log"
)

func init() {
	CurrentDb = DbSqlite
}

func ConnectDb(config *domain.Config) *sql.DB {
	if config == nil || config.Sqlite == nil {
		log.Fatal(errors.New("Publish: Database config required!"))
	}
	db, err := sql.Open("sqlite3", config.Sqlite.Path)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
