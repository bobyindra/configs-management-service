package util

import (
	"encoding/json"
	"reflect"
)

/*
This function is to compare 2 set of any data type
and converted to json data to check the values are identical or not
*/
func JsonByteEqual(a, b any) (bool, error) {
	var v1, v2 interface{}

	aJson, err := json.Marshal(a)
	if err != nil {
		return false, err
	}

	bJson, err := json.Marshal(b)
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(aJson, &v1); err != nil {
		return false, err
	}
	if err := json.Unmarshal(bJson, &v2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(v2, v1), nil
}
