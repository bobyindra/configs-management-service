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

	// Since there is no data modification in between marshal and unmarshal
	// The unmarshal will always return successfully without error
	// Error checking is not needed in this case
	json.Unmarshal(aJson, &v1)
	json.Unmarshal(bJson, &v2)

	return reflect.DeepEqual(v2, v1), nil
}
