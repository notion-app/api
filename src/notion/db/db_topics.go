package db

import (
	"notion/model"
)

func CreateTopic(t model.DbTopic) error {
	return dbmap.Insert(&t)
}
