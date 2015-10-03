package model

type School struct {
  Id string `db:"id,primary"`
  Name string `db:"name"`
  Location string `db:"location"`
}
