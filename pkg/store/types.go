package store

import "github.com/fusion-app/metric/pkg/exporter"

type Store interface{
	ListMetrics() ([]exporter.ResourceMetric, error)
	AddMetric(metric exporter.ResourceMetric) error
	UpdateMetric(name string, new exporter.ResourceMetric) (exporter.ResourceMetric, error)
}
