package metrics

import (
	"context"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

type WithOption func(*Metrics)

// New returns a new Metrics handler.
func New(
	ctx context.Context,
	opts ...WithOption,
) (*Metrics, error) {
	metrics := &Metrics{}
	for _, opt := range opts {
		opt(metrics)
	}
	if metrics.exporter == nil {
		exp, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
		if err != nil {
			return nil, err
		}
		metrics.exporter = exp
	}
	if metrics.reader == nil {
		reader := metric.NewPeriodicReader(metrics.exporter)
		metrics.reader = reader
	}
	if metrics.mp == nil {
		mp := metric.NewMeterProvider(
			metric.WithView(Views...),
			metric.WithReader(metrics.reader),
		)
		metrics.mp = mp
	}
	return metrics, nil
}

// WithMeterProvider sets the Metrics handler's MeterProvider.
func WithMeterProvider(mp *metric.MeterProvider) WithOption {
	return func(m *Metrics) {
		m.mp = mp
	}
}

// WithExporter sets the Metrics handler's Exporter.
func WithExporter(exp metric.Exporter) WithOption {
	return func(m *Metrics) {
		m.exporter = exp
	}
}

// WithReader sets the Metrics handler's Reader.
func WithReader(reader metric.Reader) WithOption {
	return func(m *Metrics) {
		m.reader = reader
	}
}
