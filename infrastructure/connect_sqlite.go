// +build sqlite

package infrastructure

import (
	"database/sql"
	"errors"
	"log"

	"github.com/kendellfab/publish/domain"
	_ "github.com/mxk/go-sqlite/sqlite3"
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
