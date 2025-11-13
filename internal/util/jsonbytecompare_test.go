package util_test

import (
	"testing"

	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestUtil_JsonByteEqual(t *testing.T) {
	t.Parallel()

	t.Run("Compare jsondata success return correct result", func(t *testing.T) {
		t.Parallel()

		// Given couple of []map[string] data
		dataToCompare := []map[string]interface{}{
			{"name": "payment-config", "enabled": true},
			{"enabled": true, "name": "payment-config"},
			{"enabled": false, "name": "payment-config"},
			{"name": "payment-config", "enabled": true, "description": "lorem ipsum"},
		}

		// When call JsonByteEqual function with identical data set
		res, _ := util.JsonByteEqual(dataToCompare[0], dataToCompare[1])

		// Then return true
		assert.True(t, res)

		// When call JsonByteEqual function with non identical data set
		res, _ = util.JsonByteEqual(dataToCompare[0], dataToCompare[2])

		// Then return false
		assert.False(t, res)

		// When call JsonByteEqual function with non identical data set
		res, _ = util.JsonByteEqual(dataToCompare[0], dataToCompare[3])

		// Then return false
		assert.False(t, res)

		// Given String data
		string1 := "Test"
		string2 := "Test"
		string3 := "Case"

		// When compare the same string value
		res, _ = util.JsonByteEqual(string1, string2)

		// Then return true
		assert.True(t, res)

		// When compare the different string value
		res, _ = util.JsonByteEqual(string1, string3)

		// Then return false
		assert.False(t, res)

		// When compare the different string value
		res, _ = util.JsonByteEqual(dataToCompare[0], string3)

		// Then return false
		assert.False(t, res)
	})

	t.Run("Given unsupported datatype for json on first data return marshal error", func(t *testing.T) {
		t.Parallel()

		// Given invalid first data
		a := make(chan int)
		b := "data"

		// When
		valid, err := util.JsonByteEqual(a, b)

		// Then
		assert.False(t, valid)
		assert.Error(t, err)
	})

	t.Run("Given unsupported datatype for json on second data return marshal error", func(t *testing.T) {
		t.Parallel()

		// Given invalid second data
		a := "data"
		b := make(chan int)

		// When
		valid, err := util.JsonByteEqual(a, b)

		// Then
		assert.False(t, valid)
		assert.Error(t, err)
	})
}
