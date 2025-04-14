package errs

import "net/http"

type ApiError struct {
	StatusCode    int    `json:"-"`
	Message       string `json:"message"`
	Cause         string `json:"cause,omitempty"`
	ErrorInstance error  `json:"-"`
}

func (a ApiError) Error() string {
	if a.Cause != "" {
		return a.Cause
	}

	return a.Message
}

// Creates a new ApiError instance with the given message, error and status code.
// Parameters:
//   - message: The error message to include
//   - err: The underlying error (nullable)
//   - statusCode: The HTTP status code to associate with this error
//
// Returns:
//   - ApiError: A new ApiError instance with the provided details
func NewApiError(message string, err error, statusCode int) ApiError {
	if err == nil {
		return ApiError{Message: message, StatusCode: statusCode}
	}

	return ApiError{Message: message, Cause: err.Error(), StatusCode: statusCode}
}

// Returns a new api error with status bad request
func BadRequestError(message string) ApiError {
	return ApiError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

// Returns a new api error with status internal server error
func InternalError(err error) ApiError {
	return ApiError{
		Message:    "Internal Server Error",
		StatusCode: http.StatusInternalServerError,
		Cause:      err.Error(),
	}
}

// Returns a new api error with status not found
func NotFoundError(message string) ApiError {
	return ApiError{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}
