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
	ErrConfigSchemaNotFound = &ErrorDetail{
		Message:  "Configuration schema not found",
		Code:     "CONFIG_SCHEMA_NOT_FOUND",
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

	ErrNoChangesFound = &ErrorDetail{
		Message:  "Configs values is the same with current configs values",
		Code:     "NO_CHANGES_FOUND",
		HttpCode: http.StatusBadRequest,
	}

	ErrRollbackNotAllowed = &ErrorDetail{
		Message:  "Cannot rollback to the same version as current",
		Code:     "ROLLBACK_NOT_ALLOWED",
		HttpCode: http.StatusBadRequest,
	}

	ErrInvalidLogin = &ErrorDetail{
		Message:  "Inccorect login info",
		Code:     "INVALID_LOGIN",
		HttpCode: http.StatusBadRequest,
	}

	ErrForbidden = &ErrorDetail{
		Message:  "You don't have permission",
		Code:     "FORBIDDEN",
		HttpCode: http.StatusForbidden,
	}

	ErrInvalidRequestParameters = &ErrorDetail{
		Message:  "Invalid request parameters",
		Code:     "INVALID_REQUEST_PARAMETERS",
		HttpCode: http.StatusBadRequest,
	}
)

func NewError(code string, message string, httpCode int) *ErrorDetail {
	return &ErrorDetail{
		Code:     code,
		Message:  message,
		HttpCode: httpCode,
	}
}

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
