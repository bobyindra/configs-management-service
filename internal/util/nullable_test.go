package util_test

import (
	"testing"
	"time"

	"github.com/bobyindra/configs-management-service/internal/testutil"
	"github.com/bobyindra/configs-management-service/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestUtil_GeneralNullable(t *testing.T) {
	t.Parallel()

	t.Run("GeneralNullable success - int", func(t *testing.T) {
		t.Parallel()

		res := util.GeneralNullable(10)
		assert.Equal(t, *res, 10)
	})

	t.Run("GeneralNullable success - string", func(t *testing.T) {
		t.Parallel()

		res := util.GeneralNullable("abc")
		assert.Equal(t, *res, "abc")
	})

	t.Run("GeneralNullable success - time", func(t *testing.T) {
		t.Parallel()

		res := util.GeneralNullable(testutil.CreatedAt)
		assert.Equal(t, *res, testutil.CreatedAt)
	})

}

func TestUtil_GeneralNullableCollection(t *testing.T) {
	t.Parallel()

	t.Run("GeneralNullableCollection success - int", func(t *testing.T) {
		t.Parallel()

		res := util.GeneralNullableCollection(1, 2, 3)
		assert.Equal(t, *res[0], 1)
		assert.Equal(t, *res[1], 2)
		assert.Equal(t, *res[2], 3)
	})

	t.Run("GeneralNullableCollection success - string", func(t *testing.T) {
		t.Parallel()

		res := util.GeneralNullableCollection("a", "b")
		assert.Equal(t, *res[0], "a")
		assert.Equal(t, *res[1], "b")
	})
}

func TestUtil_GeneralNonNullable(t *testing.T) {
	t.Parallel()

	t.Run("GeneralNonNullable success - *int", func(t *testing.T) {
		t.Parallel()

		var a int = 10

		res := util.GeneralNonNullable(&a)
		assert.Equal(t, res, 10)
	})

	t.Run("GeneralNonNullable success - *string", func(t *testing.T) {
		t.Parallel()

		var a string = "abc"

		res := util.GeneralNonNullable(&a)
		assert.Equal(t, res, "abc")
	})

	t.Run("GeneralNonNullable success - *time", func(t *testing.T) {
		t.Parallel()

		var a time.Time = testutil.CreatedAt

		res := util.GeneralNonNullable(&a)
		assert.Equal(t, res, testutil.CreatedAt)
	})

	t.Run("GeneralNonNullable success - empty int", func(t *testing.T) {
		t.Parallel()

		res := util.GeneralNonNullable[int](nil)
		assert.Equal(t, res, 0)
	})

	t.Run("GeneralNonNullable success - empty string", func(t *testing.T) {
		t.Parallel()

		res := util.GeneralNonNullable[string](nil)
		assert.Equal(t, res, "")
	})

}
