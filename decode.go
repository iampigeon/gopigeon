package iampigeon

import (
	"encoding/json"
)

// Decode return an error if the map cannot be decoded into Given struct struct
// In case of r is not a pointer, returns error
func Decode(m map[string]interface{}, r interface{}) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	return nil
}
