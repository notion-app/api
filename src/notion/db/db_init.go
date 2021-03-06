package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	"notion/config"
	"notion/log"
	"notion/model"
	"os"
)

var (
	dbmap *gorp.DbMap
)

// Init creates a connection to the database
func Init() {
	log.Info("Establishing a connection to the database")
	db, err := sql.Open("postgres", config.PostgresURL())
	if log.Error(err) {
		os.Exit(1)
	}
	err = db.Ping()
	if log.Error(err) {
		os.Exit(1)
	}
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(model.DbSchool{}, "schools").SetKeys(false, "Id")
	dbmap.AddTableWithName(model.DbCourse{}, "courses").SetKeys(false, "Id")
	dbmap.AddTableWithName(model.DbCourseSection{}, "sections").SetKeys(false, "Id")
	dbmap.AddTableWithName(model.DbUser{}, "users").SetKeys(false, "Id")
	dbmap.AddTableWithName(model.DbSchoolRequest{}, "school_requests").SetKeys(false, "Id")
	dbmap.AddTableWithName(model.DbNotebook{}, "notebooks").SetKeys(false, "Id")
	dbmap.AddTableWithName(model.DbTopic{}, "topics").SetKeys(false, "Id")
	dbmap.AddTableWithName(model.DbNote{}, "notes").SetKeys(false, "Id")
	dbmap.AddTableWithName(model.DbSubscription{}, "subscriptions").SetKeys(false, "UserId", "NotebookId")
}
