package db

import (
	"net/http"
	"notion/errors"
	"notion/model"
)

func GetNotebookById(id string) (model.DbNotebook, error) {
	var nb model.DbNotebook
	in, err := GenericGetOne(&nb, "select * from notebooks where id=$1", id)
	if !in {
		return nb, errors.NewHttp(http.StatusNotFound, "The notebook requested was not found")
	}
	return nb, err
}
