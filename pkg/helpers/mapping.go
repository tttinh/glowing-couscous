package helpers

import (
	"encoding/json"
)

// Copy value as json
func Copy(toValue interface{}, fromValue interface{}) error {
	data, err := json.Marshal(fromValue)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, toValue)
	if err != nil {
		return err
	}

	return nil
}
