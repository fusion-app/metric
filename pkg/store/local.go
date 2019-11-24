package store

import "github.com/fusion-app/metric/pkg/exporter"

var metrics []exporter.ResourceMetric

type LocalStore struct{}

func (s *LocalStore) ListMetrics() ([]exporter.ResourceMetric, error) {
	return metrics, nil
}

func (s *LocalStore) AddMetric(metric exporter.ResourceMetric) error {
	return nil
}

func (s *LocalStore) UpdateMetric(name string, new exporter.ResourceMetric) (exporter.ResourceMetric, error) {
	return exporter.ResourceMetric{}, nil
}