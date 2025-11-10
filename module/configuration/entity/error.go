package entity

import (
	"net/http"
)

type ErrorDetail struct {
	Message  string `json:"message"`
	Code     string `json:"code"`
	Field    string `json:"field,omitempty"`
	HttpCode int    `json:"-"`
}

func (e *ErrorDetail) Error() string {
	return e.Message
}

var (
	ErrConfigNotFound = &ErrorDetail{
		Message:  "Configuration not found",
		Code:     "CONFIG_NOT_FOUND",
		HttpCode: http.StatusNotFound,
	}

	ErrInvalidSchema = &ErrorDetail{
		Message:  "Invalid Schema",
		Code:     "INVALID_SCHEMA",
		HttpCode: http.StatusBadRequest,
	}

	ErrConfigAlreadyExists = &ErrorDetail{
		Message:  "Config schema already exists",
		Code:     "CONFIG_ALREADY_EXISTS",
		HttpCode: http.StatusConflict,
	}

	ErrNoChangesDetected = &ErrorDetail{
		Message:  "No changes detected",
		Code:     "NO_CHANGES_DETECTED",
		HttpCode: http.StatusBadRequest,
	}

	ErrInvalidConfigValues = &ErrorDetail{
		Message:  "Invalid config values",
		Code:     "INVALID_CONFIG_VALUES",
		HttpCode: http.StatusBadRequest,
	}

	ErrInvalidRequestParameters = &ErrorDetail{
		Message:  "Invalid request parameters",
		Code:     "INVALID_REQUEST_PARAMETERS",
		HttpCode: http.StatusBadRequest,
	}
)

func WrapError(err error) *ErrorDetail {
	return &ErrorDetail{
		Message:  err.Error(),
		Code:     "INTERNAL_ERROR",
		HttpCode: http.StatusInternalServerError,
	}
}

func ErrEmptyField(fieldName string) *ErrorDetail {
	return &ErrorDetail{
		Message:  fieldName + " cannot be empty",
		Code:     "EMPTY_FIELD",
		Field:    fieldName,
		HttpCode: http.StatusBadRequest,
	}
}

func ErrNotFound(fieldName string) *ErrorDetail {
	return &ErrorDetail{
		Message:  fieldName + " not found",
		Code:     "NOT_FOUND",
		HttpCode: http.StatusNotFound,
	}
}

func ErrConfigVersionNotFound(fieldName string) *ErrorDetail {
	return &ErrorDetail{
		Message:  "Config Version " + fieldName + " not found",
		Code:     "CONFIG_VERSION_NOT_FOUND",
		HttpCode: http.StatusNotFound,
	}
}
