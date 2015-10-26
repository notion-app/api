
package db

import (
  "database/sql"
  "notion/log"
)

func GenericGetOne(model interface{}, sqls string, args ...interface{}) (bool, error) {
	err := dbmap.SelectOne(model, sqls, args...)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return false, nil
		default:
			log.Error(err)
			return false, err
		}
	}
	return true, nil
}

// model passed in is assumed to be a list
// because of this, a boolean is not returned; you can just check the length of the list
func GenericGet(model interface{}, sqls string, args ...interface{}) error {
  _, err := dbmap.Select(model, sqls, args...)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil
		default:
			log.Error(err)
			return err
		}
	}
	return nil
}
