package util_test

import (
	"testing"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestUtil_JsonByteEqual(t *testing.T) {
	t.Parallel()

	t.Run("Compare identical jsondata success", func(t *testing.T) {
		t.Parallel()

		// Try to compare map[string]interface{}
		dataToCompare := []map[string]interface{}{
			{"name": "payment-config", "enabled": true},
			{"enabled": true, "name": "payment-config"},
			{"enabled": false, "name": "payment-config"},
			{"name": "payment-config", "enabled": true, "description": "lorem ipsum"},
		}

		res, _ := util.JsonByteEqual(dataToCompare[0], dataToCompare[1])
		assert.True(t, res)

		res, _ = util.JsonByteEqual(dataToCompare[0], dataToCompare[2])
		assert.False(t, res)

		res, _ = util.JsonByteEqual(dataToCompare[0], dataToCompare[3])
		assert.False(t, res)

		string1 := "Test"
		string2 := "Test"
		string3 := "Case"

		res, _ = util.JsonByteEqual(string1, string2)
		assert.True(t, res)

		res, _ = util.JsonByteEqual(string1, string3)
		assert.False(t, res)

		res, _ = util.JsonByteEqual(dataToCompare[0], string3)
		assert.False(t, res)
	})
}
