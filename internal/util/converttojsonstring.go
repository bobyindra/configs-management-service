package util

import "encoding/json"

func ConvertAnyValueToJsonString(value any) (*string, error) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	jsonString := string(jsonValue)

	return &jsonString, nil
}
