package usecase

import (
	"context"
	"math"
	"math/rand"

	"github.com/hackThacker/OWASP-AttackForge/backend/domain"
)

type metricUsecase struct {
	toolRepo domain.ToolRepository
}

func NewMetricUsecase(toolRepo domain.ToolRepository) domain.MetricUsecase {
	return &metricUsecase{toolRepo: toolRepo}
}

func (u *metricUsecase) GetMetrics(ctx context.Context) (*domain.Metric, error) {
	tools, err := u.toolRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	total := len(tools)
	running := 0
	for _, t := range tools {
		if !t.Stopped {
			running++
		}
	}

	health := 0
	if total > 0 {
		health = (running * 100) / total
	}

	// Calculate realistic CPU and memory footprints based on container status
	var cpu float64
	var mem float64
	if running > 0 {
		// Base CPU load: 10% + 4.2% per running target lab + random jitter
		cpu = 10.0 + (float64(running) * 4.2) + (rand.Float64() * 2.0)
		cpu = math.Round(cpu*10) / 10

		// Base Memory load: 3.5 GB + 0.38 GB per running target lab + random jitter
		mem = 3.5 + (float64(running) * 0.38) + (rand.Float64() * 0.15)
		mem = math.Round(mem*10) / 10
	} else {
		cpu = 0.5
		mem = 0.2
	}

	return &domain.Metric{
		TotalContainers:   total,
		RunningContainers: running,
		HealthPercentage:  health,
		CPUUsage:          cpu,
		MemoryUsage:       mem,
	}, nil
}
