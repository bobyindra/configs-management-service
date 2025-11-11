package entity_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/stretchr/testify/assert"
)

func TestEntity_Error(t *testing.T) {
	t.Parallel()

	t.Run("NewError success return correct error data", func(t *testing.T) {
		t.Parallel()

		// Given error data set
		code := "INVALID_REQUEST"
		message := "Invalid request"
		httpCode := http.StatusBadRequest

		// When call NewError function
		res := entity.NewError(code, message, httpCode)

		// Then all data should return correctly
		assert.Equal(t, code, res.Code)
		assert.Equal(t, message, res.Message)
		assert.Equal(t, httpCode, res.HttpCode)
	})

	t.Run("WrapError success return correct error data", func(t *testing.T) {
		t.Parallel()

		// Given error data set
		err := errors.New("Some Error")
		code := "INTERNAL_ERROR"
		httpCode := http.StatusInternalServerError

		// When call WrapError function
		res := entity.WrapError(err)

		// Then all data should return correctly
		assert.Equal(t, code, res.Code)
		assert.Equal(t, err.Error(), res.Message)
		assert.Equal(t, httpCode, res.HttpCode)
	})

	t.Run("ErrEmptyField success return correct error data", func(t *testing.T) {
		t.Parallel()

		// Given error data set
		code := "EMPTY_FIELD"
		field := "name"
		httpCode := http.StatusBadRequest

		// When call ErrEmptyField function
		res := entity.ErrEmptyField(field)

		// Then all data should return correctly
		assert.Equal(t, code, res.Code)
		assert.Equal(t, field+" cannot be empty", res.Message)
		assert.Equal(t, httpCode, res.HttpCode)
	})

	t.Run("ErrNotFound success return correct error data", func(t *testing.T) {
		t.Parallel()

		// Given error data set
		code := "NOT_FOUND"
		field := "version"
		httpCode := http.StatusNotFound

		// When call ErrEmptyField function
		res := entity.ErrNotFound(field)

		// Then all data should return correctly
		assert.Equal(t, code, res.Code)
		assert.Equal(t, field+" not found", res.Message)
		assert.Equal(t, httpCode, res.HttpCode)
	})

	t.Run("ErrConfigVersionNotFound success return correct error data", func(t *testing.T) {
		t.Parallel()

		// Given error data set
		code := "CONFIG_VERSION_NOT_FOUND"
		field := "test-config"
		httpCode := http.StatusNotFound

		// When call ErrEmptyField function
		res := entity.ErrConfigVersionNotFound(field)

		// Then all data should return correctly
		assert.Equal(t, code, res.Code)
		assert.Equal(t, "Config Version "+field+" not found", res.Message)
		assert.Equal(t, httpCode, res.HttpCode)
	})

	t.Run("Error success return correct error message", func(t *testing.T) {
		t.Parallel()

		// Given error data set
		message := "Error Found"
		data := &entity.ErrorDetail{
			Message: message,
		}

		// When call ErrEmptyField function
		res := data.Error()

		// Then all data should return correctly
		assert.Equal(t, message, res)
	})
}
