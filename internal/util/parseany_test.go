package util_test

import (
	"testing"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestUtil_ParseAny(t *testing.T) {
	t.Parallel()

	t.Run("ParseAny jsonstring to map[string]interface{} success", func(t *testing.T) {
		t.Parallel()
		input := "{\"name\":\"payment-config\"}"
		expected := map[string]interface{}{"name": "payment-config"}

		res := util.ParseAny(input)
		assert.Equal(t, expected, res)

	})

	t.Run("ParseAny stringboolean to boolean success", func(t *testing.T) {
		t.Parallel()
		input := "true"
		expected := true

		res := util.ParseAny(input)
		assert.Equal(t, expected, res)
	})

	t.Run("ParseAny unsupported json data return error", func(t *testing.T) {
		t.Parallel()
		input := make(chan int)
		expected := input

		res := util.ParseAny(input)
		assert.Equal(t, expected, res)
	})
}
