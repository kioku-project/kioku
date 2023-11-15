package helper

import (
	"context"
	"io"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func getExporter(ctx context.Context) (trace.SpanExporter, error) {
	if os.Getenv("TRACING_ENABLED") != "true" {
		return stdouttrace.New(stdouttrace.WithWriter(io.Discard))
	}

	tracingUrl := os.Getenv("TRACING_COLLECTOR")
	if tracingUrl == "" {
		tracingUrl = "simple-prod-collector.observability.svc.cluster.local:4318"
	}
	return otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(tracingUrl),
		otlptracehttp.WithInsecure(),
	)

}

func SetupTracing(ctx context.Context, serviceName string) (*trace.TracerProvider, error) {

	exporter, err := getExporter(ctx)
	if err != nil {
		return nil, err
	}

	// labels/tags/resources that are common to all traces.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
	)

	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource),
		// set the sampling rate based on the parent span to 60%
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(0.6))),
	)

	otel.SetTracerProvider(provider)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{}, // W3C Trace Context format; https://www.w3.org/TR/trace-context/
			propagation.Baggage{},
		),
	)

	return provider, nil
}
