
package db

import (
  "notion/model"
)

func GetSchool(schoolId string) (bool, model.DbSchool, error) {
  var school model.DbSchool
  in, err := GenericGetOne(&school, "select * from schools where id=$1", schoolId)
  return in, school, err
}

func GetAllSchools() ([]model.DbSchool, error) {
  schools := []model.DbSchool{}
  err := GenericGet(&schools, `select * from schools`)
  return schools, err
}

func GetSectionByNotebookId(notebookId string) (model.DbCourseSection, error) {
  var section model.DbCourseSection
  _, err := GenericGetOne(&section, `select * from sections where notebook_id=$1`, notebookId)
  return section, err
}

func GetCoursesForSchool(schoolId string) ([]model.DbCourse, error) {
  courses := []model.DbCourse{}
  err := GenericGet(&courses, `select * from courses where school_id=$1`, schoolId)
  return courses, err
}

func GetCourseByCourseId(courseId string) (model.DbCourse, error) {
  var course model.DbCourse
  _, err := GenericGetOne(&course, `select * from courses where id=$1`, courseId)
  return course, err
}

func GetSectionsForCourse(courseId string) ([]model.DbCourseSection, error) {
  sections := []model.DbCourseSection{}
  err := GenericGet(&sections, `select * from sections where course_id=$1`, courseId)
  return sections, err
}

func CreateSchoolRequest(s model.DbSchoolRequest) error {
  return dbmap.Insert(&s)
}
