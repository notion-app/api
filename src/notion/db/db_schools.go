
package db

import (
  "notion/model"
)

func GetAllSchools() ([]model.DbSchool, error) {
  schools := []model.DbSchool{}
  err := GenericGet(&schools, `select * from schools`)
  return schools, err
}
