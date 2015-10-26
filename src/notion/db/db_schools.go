
package db

import (
  "notion/model"
)

func GetAllSchools() ([]model.DbSchool, error) {
  schools := []model.DbSchool{}
  err := GenericGet(&schools, `select * from schools`)
  return schools, err
}

func GetCoursesForSchool(schoolId string) ([]model.DbCourse, error) {
  courses := []model.DbCourse{}
  err := GenericGet(&courses, `select * from courses where school_id=$1`, schoolId)
  return courses, err
}
