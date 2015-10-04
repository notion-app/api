package model

type School struct {
  Id string `db:"id,primary" json:"id"`
  Name string `db:"name" json:"name"`
  Location string `db:"location" json:"location"`
}

type Course struct {
  Id string `db:"id,primary" json:"id"`
  SchoolId string `db:"school_id" json:"school_id"`
  Name string `db:"name" json:"name"`
  Number string `db:"number" json:"number"`
}

type CourseSection struct {
  Id string `db:"id,primary" json:"id"`
  CourseId string `db:"course_id" json:"course_id"`
  Crn string `db:"crn" json:"crn"`
  Professor string `db:"professor" json:"professor"`
  Year string `db:"year" json:"year"`
  Semester string `db:"semester" json:"semester"`
  Time string `db:"time" json:"time"`
  Verified bool `db:"verified" json:"verified"`
}
