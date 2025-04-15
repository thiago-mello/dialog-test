package errs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

// ErrorHandler returns an echo.HTTPErrorHandler that handles API and unknown errors.
// It checks if the response is already committed and processes different error types accordingly.
func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		switch e := err.(type) {
		case ApiError:
			handleApiError(e, c)
		case *echo.HTTPError:
			handleHttpError(e, c)
		default:
			handleUnknownError(e, c)
		}
	}
}

// handleApiError processes API-specific errors by logging a warning and returning a JSON response.
// If debug mode is enabled, it includes the full error details in the response.
// Parameters:
//   - err: The API error that occurred
//   - c: The echo.Context for the request
func handleApiError(err ApiError, c echo.Context) {
	c.Logger().Warn(err.Error())

	message := err.Message
	var details string
	if c.Echo().Debug {
		details = err.Error()
	}

	c.JSON(err.StatusCode, ErrorResponse{
		Message: message,
		Details: details,
	})
}

// handleHttpError processes HTTP-specific errors by logging a warning and returning a JSON response.
// If the error message is not a string, it uses the standard HTTP status text.
// If debug mode is enabled, it includes the full error details in the response.
// Parameters:
//   - err: The HTTP error that occurred
//   - c: The echo.Context for the request
func handleHttpError(err *echo.HTTPError, c echo.Context) {
	c.Logger().Warn(err.Error())

	message, ok := err.Message.(string)
	if !ok {
		message = http.StatusText(err.Code)
	}

	var details string
	if c.Echo().Debug {
		details = err.Error()
	}

	c.JSON(err.Code, ErrorResponse{
		Message: message,
		Details: details,
	})
}

// handleUnknownError processes unexpected errors by logging them and returning a generic error response.
// If debug mode is enabled, it includes the full error details in the response.
// Parameters:
//   - err: The unknown error that occurred
//   - c: The echo.Context for the request
func handleUnknownError(err error, c echo.Context) {
	c.Logger().Error(err)

	var details string
	if c.Echo().Debug {
		details = err.Error()
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Message: "Internal server error",
		Details: details,
	})
}
