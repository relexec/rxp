package metrics

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

// SumResourceMetrics returns the int64 sum of the named instrument within the
// supplied metricsdata.ResourceMetrics.
//
// If the named instrument does not exist in the supplied ResourceMetrics,
// returns an error.
//
// If the named instrument exists in the supplied ResourceMetrics but is not an
// int64 sum, returns an error.
func SumResourceMetrics(
	data *metricdata.ResourceMetrics,
	instrument string,
) (int64, error) {
	for _, sm := range data.ScopeMetrics {
		m, ok := lo.Find(sm.Metrics, func(subject metricdata.Metrics) bool {
			return strings.EqualFold(instrument, subject.Name)
		})
		sum, ok := m.Data.(metricdata.Sum[int64])
		if !ok {
			return -1, fmt.Errorf(
				"instrument %q is a %T, not a metricsdata.Sum[int64]",
				instrument, m.Data,
			)
		}
		return lo.SumBy(
			sum.DataPoints,
			func(dp metricdata.DataPoint[int64]) int64 {
				return dp.Value
			},
		), nil
	}
	return -1, fmt.Errorf(
		"instrument %q not in supplied ResourceMetrics", instrument,
	)
}

// InResourceMetrics returns true if the named instrument is within
// the supplied metricdata.ResourceMetrics.
func InResourceMetrics(
	data *metricdata.ResourceMetrics,
	instrument string,
) bool {
	for _, sm := range data.ScopeMetrics {
		if lo.ContainsBy(
			sm.Metrics,
			func(subject metricdata.Metrics) bool {
				return strings.EqualFold(instrument, subject.Name)
			},
		) {
			return true
		}
	}
	return false
}
