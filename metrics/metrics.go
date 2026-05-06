package metrics

import (
	"go.opentelemetry.io/otel/sdk/metric"
)

// Metrics handles OTEL Meters for the rxp store
type Metrics struct {
	// mp is the OTEL metric.MeterProvider
	mp *metric.MeterProvider
	// exporter is the OTEL sdkmetric.Exporter
	exporter metric.Exporter
	// reader is the OTEL sdkmetric.Reader
	reader metric.Reader
}

// MeterProvider returns the OTEL metric.MeterProvider
func (m Metrics) MeterProvider() *metric.MeterProvider {
	return m.mp
}

// Exporter returns the OTEL sdkmetric.Exporter.
func (m Metrics) Exporter() metric.Exporter {
	return m.exporter
}

// Reader returns the OTEL sdkmetric.Reader.
func (m Metrics) Reader() metric.Reader {
	return m.reader
}
