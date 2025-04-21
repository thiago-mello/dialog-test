package instrumentation

import (
	"context"

	"github.com/leandro-andrade-candido/api-go/src/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// InitTracer initializes and configures an OpenTelemetry tracer provider
// It sets up OTLP HTTP exporter, resource attributes, and trace propagation
//
// Parameters:
//   - ctx: Context for the initialization process
//
// Returns:
//   - *trace.TracerProvider: Configured tracer provider instance
//   - error: Any error that occurred during initialization
//
// The function will:
// - Get the OTLP traces endpoint from config
// - Create an OTLP HTTP exporter
// - Configure resource attributes including service name and debug flag
// - Create and configure the tracer provider
// - Set up trace context and baggage propagation
// - Return nil provider if no endpoint configured
func InitTracer(ctx context.Context) (*trace.TracerProvider, error) {
	otlpTracesEndpoint := config.GetString("otel.traces.otlp.endpoint")
	if otlpTracesEndpoint == "" {
		return nil, nil
	}

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(otlpTracesEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("dialog-api"),
			attribute.Bool("debug", config.GetBoolean("server.debug")),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return tp, nil
}
