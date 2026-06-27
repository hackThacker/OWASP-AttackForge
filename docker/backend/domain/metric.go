package domain

import "context"

type Metric struct {
	TotalContainers   int     `json:"totalContainers"`
	RunningContainers int     `json:"runningContainers"`
	HealthPercentage  int     `json:"healthPercentage"`
	CPUUsage          float64 `json:"cpuUsage"`
	MemoryUsage       float64 `json:"memoryUsage"`
}

type MetricUsecase interface {
	GetMetrics(ctx context.Context) (*Metric, error)
}
