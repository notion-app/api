
package db

import (
  "notion/model"
)

// Gets the list of all schools in the database
func GetAllSchools() ([]model.DbSchool, error) {
  var schools []model.DbSchool
  schoolsG, err := GenericGetMultiple("schools", "", "", &schools)
  return *schoolsG.(*[]model.DbSchool), err
}
