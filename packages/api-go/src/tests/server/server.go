package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/validation"
)

var e *echo.Echo = nil

func GetServer() *echo.Echo {
	if e == nil {
		e = echo.New()
		e.Debug = true
		configErrorHandler()
		configureServerInputValidation()
	}

	return e
}

// configures default error handler
func configErrorHandler() {
	e.HTTPErrorHandler = errs.ErrorHandler()
}

// configures input validator
func configureServerInputValidation() {
	e.Validator = &validation.Validator{Validator: validator.New()}
}
