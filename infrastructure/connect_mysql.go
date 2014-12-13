// +build !sqlite

package infrastructure

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kendellfab/publish/domain"
	"log"
)

func init() {
	CurrentDb = DbMysql
}

func ConnectDb(config *domain.Config) *sql.DB {
	if config == nil || config.Mysql == nil {
		log.Fatal(errors.New("Publish: Database config required!"))
	}
	connection := config.Mysql.GetConnectionString()
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	if config.Mysql.MaxIdle != 0 {
		db.SetMaxIdleConns(config.Mysql.MaxIdle)
	}
	if config.Mysql.MaxOpen != 0 {
		db.SetMaxOpenConns(config.Mysql.MaxOpen)
	}
	return db
}
