package database

import (
	"database/sql"
	"log"

	"github.com/google/wire"
)

// Set is a Wire provider set that produces a *sql.DB.
var Set = wire.NewSet(
	Open,
)

// DataSource specifies how to connect to a source database.
type DataSource struct {
	DriverName     string
	DataSourceName string
}

// Open opens a connection to a SQL database.
func Open(src DataSource) (*sql.DB, func(), error) {
	db, err := sql.Open(src.DriverName, src.DataSourceName)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}
	return db, cleanup, nil
}
