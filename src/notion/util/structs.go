package util

import (
	"encoding/json"
)

func FillStruct(s interface{}, m map[string]interface{}) error {
	result, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(result, s)
}
