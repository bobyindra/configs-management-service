package util

import (
	"encoding/json"
	"fmt"
)

func ParseAny(anyType any) any {
	var parsed any
	if err := json.Unmarshal([]byte(fmt.Sprint(anyType)), &parsed); err != nil {
		parsed = anyType
	}
	return parsed
}
