package mysql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"

	"github.com/terashi58/wire-example/base/database"
)

// Set is a Wire provider set that produces a *sql.DB with MySQL driver.
var Set = wire.NewSet(
	database.Set,
	NewDataSource,
)

// NewDataSource generates a DataSource for MySQL.
func NewDataSource(cfg *mysql.Config) database.DataSource {
	return database.DataSource{
		DriverName:     "mysql",
		DataSourceName: cfg.FormatDSN(),
	}
}
