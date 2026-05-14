package metrics

import (
	"go.opentelemetry.io/otel/metric"

	"github.com/relexec/rxp/types"
)

const (
	ProviderName = "rxp"

	InstrumentNameReadRequest = "read.request"
	InstrumentDescReadRequest = "Number of read operations. Labels: 'error.code', 'type'."

	InstrumentNameReadDuration = "read.duration"
	InstrumentDescReadDuration = "Histogram of read operation duration in seconds. Labels: 'error.code', 'type'."

	InstrumentNameWriteRequest = "write.request"
	InstrumentDescWriteRequest = "Number of write operations. Labels: 'error.code', 'type'."

	InstrumentNameWriteDuration = "write.duration"
	InstrumentDescWriteDuration = "Histogram of write operation duration in seconds. Labels: 'error.code', 'type'."

	InstrumentNameListRequest = "list.request"
	InstrumentDescListRequest = "Number of list operations. Labels: 'error.code', 'type'."

	InstrumentNameListDuration = "list.duration"
	InstrumentDescListDuration = "Histogram of list operation duration in seconds. Labels: 'error.code', 'type'."
)

var (
	InstrumentReadRequest   metric.Int64Counter
	InstrumentReadDuration  metric.Float64Histogram
	InstrumentWriteRequest  metric.Int64Counter
	InstrumentWriteDuration metric.Float64Histogram
	InstrumentListRequest   metric.Int64Counter
	InstrumentListDuration  metric.Float64Histogram
)

// Init initializes the rxp metrics for the supplied Metrics handler.
func Init(metrics types.Metrics) error {
	var err error
	p := metrics.MeterProvider()
	m := p.Meter(ProviderName)

	InstrumentReadRequest, err = m.Int64Counter(
		InstrumentNameReadRequest,
		metric.WithDescription(InstrumentDescReadRequest),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		return err
	}

	InstrumentWriteRequest, err = m.Int64Counter(
		InstrumentNameWriteRequest,
		metric.WithDescription(InstrumentDescWriteRequest),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		return err
	}

	InstrumentReadDuration, err = m.Float64Histogram(
		InstrumentNameReadDuration,
		metric.WithDescription(InstrumentDescReadDuration),
		metric.WithUnit("{seconds}"),
	)
	if err != nil {
		return err
	}

	InstrumentWriteDuration, err = m.Float64Histogram(
		InstrumentNameWriteDuration,
		metric.WithDescription(InstrumentDescWriteDuration),
		metric.WithUnit("{seconds"),
	)
	if err != nil {
		return err
	}

	InstrumentListRequest, err = m.Int64Counter(
		InstrumentNameListRequest,
		metric.WithDescription(InstrumentDescListRequest),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		return err
	}

	InstrumentListDuration, err = m.Float64Histogram(
		InstrumentNameListDuration,
		metric.WithDescription(InstrumentDescListDuration),
		metric.WithUnit("{seconds}"),
	)
	if err != nil {
		return err
	}
	return nil
}
