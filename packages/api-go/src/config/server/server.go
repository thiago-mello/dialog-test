package server

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/config/instrumentation"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/validation"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/sdk/trace"
)

var e *echo.Echo
var tracerProvider *trace.TracerProvider

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

// configures opentelemetry tracing for echo
func configServerTelemetry() {
	tp, err := instrumentation.InitTracer(context.Background())
	if err == nil {
		e.Use(otelecho.Middleware("dialog-api"))
	}
	tracerProvider = tp
}

func GetTracerProvider() *trace.TracerProvider {
	return tracerProvider
}
