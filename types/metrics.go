package types

import (
	"go.opentelemetry.io/otel/sdk/metric"
)

// Metrics handles OTEL Meters
type Metrics interface {
	// MeterProvider returns the OTEL metric.MeterProvider
	MeterProvider() *metric.MeterProvider
	// Reader returns the OTEL sdkmetric.Reader.
	Reader() metric.Reader
	// Exporter returns the OTEL metric.
	Exporter() metric.Exporter
}
