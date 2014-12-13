// +build !sqlite

package infrastructure

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kendellfab/publish/domain"
)

func init() {
	CurrentDb = DbMysql
}

func ConnectDb(config *domain.Config) (*sql.DB, error) {
	if config == nil || config.Mysql == nil {
		return nil, errors.New("Publish: Database config required!")
	}
	connection := config.Mysql.GetConnectionString()
	return sql.Open("mysql", connection)
}
