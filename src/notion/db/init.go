package db

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"
	"notion/config"
	"notion/log"
)

var (
	db *sql.DB
)

// Init creates a connection to the database
func Init() {
	log.Info("Establishing a connection to the database")
	var err error
	db, err = sql.Open("postgres", config.PostgresURL())
	if log.Error(err) {
		os.Exit(1)
	}
	err = db.Ping()
	if log.Error(err) {
		os.Exit(1)
	}
	CreateTables(true)
}

// CreateTables does exactly that; creates the tables we need in the database.
// You can optionally specify to drop the old tables; this is primarily used for
// configuration purposes,
func CreateTables(dropOld bool) {

}
