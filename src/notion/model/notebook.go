
package model

type Notebook struct {
  Id string `db:"id,primary" json:"id"`
  SectionId string `db:"section_id" json:"section_id"`
  Name string `db:"name" json:"name"`
  Owner string `db:"owner" json:"owner"`
  Privacy string `db:"privacy" json:"privacy"`
}
