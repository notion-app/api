
package db

import (
  "notion/log"
  "notion/model"
)

// Gets the list of all schools in the database
func GetAllSchools() ([]model.DbSchool, error) {
  log.Info("Getting all schools from database")
  var schools []model.DbSchool
  _, err := dbmap.Select(&schools, "select * from schools")
  if err != nil {
    log.Error(err)
  }
  return schools, err
}
