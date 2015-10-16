package validate

import (
	"notion/errors"
	"notion/model"
)

func SchoolRequest(r model.SchoolRequestRequest) error {
	if r.Name == "" || r.Location == "" {
		return errors.BadRequest("Must supply name and location in request")
	}
	return nil
}
