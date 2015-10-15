
package db

import (
  "notion/log"
  "notion/model"
)

// Gets the list of all schools in the database
func GetAllSchools() ([]model.DbSchool, error) {
  var schools []model.DbSchool
  schoolsG, err := GenericGetMultiple("schools", "", "", &schools)
  return *schoolsG.(*[]model.DbSchool), err
}

func CreateSchoolRequest(m model.DbSchoolRequest) error {
  log.Info("Creating new school request " + m.Id)
  return dbmap.Insert(&m)
}
