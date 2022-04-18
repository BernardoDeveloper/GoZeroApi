package utils

import "encoding/json"

func Serialize(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	return string(val), nil
}
