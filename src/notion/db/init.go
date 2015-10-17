package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	"notion/config"
	"notion/errors"
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

func GenericGetOne(table string, field string, value string, model interface{}) error {
	log.Info(fmt.Sprintf("Doing single get on table %v where %v=%v", table, field, value))
	err := dbmap.SelectOne(model, fmt.Sprintf("select * from %v where %v=$1", table, field), value)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return errors.NotFound()
		default:
			log.Error(err)
			return errors.ISE()
		}
	}
	return nil
}

// model passed in is assumed to be a list
// because of this, this function will NOT return an error if nothing is found in the database.
// it would just return an empty list
func GenericGetMultiple(table string, field string, value string, model interface{}) error {
	log.Info("Doing multiple get on table %v where %v=%v", table, field, value)
	var err error
	if field == "" {
		_, err = dbmap.Select(model, fmt.Sprintf("select * from %v", table))
	} else {
		_, err = dbmap.Select(model, fmt.Sprintf("select * from %v where %v=$1", table, field), value)
	}
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil
		default:
			log.Error(err)
			return errors.ISE()
		}
	}
	return nil
}
