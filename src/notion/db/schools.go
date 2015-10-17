package db

import (
	"notion/errors"
	"notion/log"
	"notion/model"
)

// Gets the list of all schools in the database
func GetAllSchools() ([]model.DbSchool, error) {
	var schools []model.DbSchool
	err := GenericGetMultiple("schools", "", "", &schools)
	return schools, err
}

func CreateSchoolRequest(m model.DbSchoolRequest) error {
	log.Info("Creating new school request " + m.Id)
	err := dbmap.Insert(&m)
	if log.Error(err) {
		return errors.ISE()
	}
	return nil
}
