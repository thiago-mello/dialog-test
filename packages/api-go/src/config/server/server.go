package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/validation"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

var e *echo.Echo

func init() {
	e = echo.New()
	e.Debug = config.GetBoolean("server.debug")

	configErrorHandler()
	configServerTelemetry()
	configureServerInputValidation()
}

func GetServer() *echo.Echo {
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

// configures opentelemetry middleware for echo endpoints
func configServerTelemetry() {
	e.Use(otelecho.Middleware("dialog-api"))
}
