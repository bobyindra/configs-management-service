package util_test

import (
	"testing"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestUtil_ConvertToJsonString(t *testing.T) {
	t.Parallel()
	t.Run("Convert any to Json string success", func(t *testing.T) {
		t.Parallel()

		// Given data to convert to json string
		data := map[string]interface{}{"config": "test"}
		expectedResult := `{"config":"test"}`

		// When
		res, err := util.ConvertAnyValueToJsonString(data)

		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, *res)
	})

	t.Run("Convert any to Json string marshal error", func(t *testing.T) {
		t.Parallel()

		// Given data to convert to json string
		data := map[string]interface{}{"config": make(chan int)}

		// When
		res, err := util.ConvertAnyValueToJsonString(data)

		// Then
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
