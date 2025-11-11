package util

import (
	"encoding/json"
	"reflect"
)

func JsonByteEqual(a, b any) (bool, error) {
	var j, j2 interface{}

	aJson, err := json.Marshal(a)
	if err != nil {
		return false, err
	}

	bJson, err := json.Marshal(b)
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(aJson, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(bJson, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
