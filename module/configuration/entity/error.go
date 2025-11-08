package entity

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Field   string `json:"field,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

var (
	ErrConfigNotFound = &ErrorResponse{
		Message: "Configuration not found",
		Code:    "CONFIG_NOT_FOUND",
	}

	ErrSchemaDoesNotMatch = &ErrorResponse{
		Message: "Schema does not match",
		Code:    "SCHEMA_DOES_NOT_MATCH",
	}

	ErrConfigAlreadyExists = &ErrorResponse{
		Message: "Config already exists",
		Code:    "CONFIG_ALREADY_EXISTS",
	}

	ErrNoChangesDetected = &ErrorResponse{
		Message: "No changes detected",
		Code:    "NO_CHANGES_DETECTED",
	}

	ErrInvalidConfigValues = &ErrorResponse{
		Message: "Invalid config values",
		Code:    "INVALID_CONFIG_VALUES",
	}
)

func WrapError(err error) *ErrorResponse {
	return &ErrorResponse{
		Message: err.Error(),
		Code:    "INTERNAL_ERROR",
	}
}
