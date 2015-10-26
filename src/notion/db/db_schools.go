
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

func GetSectionsForCourse(courseId string) ([]model.DbCourseSection, error) {
  sections := []model.DbCourseSection{}
  err := GenericGet(&sections, `select * from sections where course_id=$1`, courseId)
  return sections, err
}

func CreateSchoolRequest(s model.DbSchoolRequest) error {
  return dbmap.Insert(&s)
}
