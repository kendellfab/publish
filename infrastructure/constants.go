package infrastructure

var CurrentDb DbType

type DbType int

const (
	DbUnknown DbType = iota
	DbSqlite
	DbMysql
)

func (t DbType) String() string {
	switch t {
	case DbSqlite:
		return "Sqlite"
	case DbMysql:
		return "MySql"
	default:
		return "Unknown"
	}
}
