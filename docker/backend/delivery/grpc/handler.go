package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/hackThacker/OWASP-AttackForge/backend/domain"
	"github.com/hackThacker/OWASP-AttackForge/backend/pb"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedAttackForgeServiceServer
	toolUC   domain.ToolUsecase
	metricUC domain.MetricUsecase
}

func RegisterServer(s *grpc.Server, toolUC domain.ToolUsecase, metricUC domain.MetricUsecase) {
	pb.RegisterAttackForgeServiceServer(s, &grpcHandler{
		toolUC:   toolUC,
		metricUC: metricUC,
	})
}

func (h *grpcHandler) GetMetrics(ctx context.Context, req *pb.MetricsRequest) (*pb.MetricsResponse, error) {
	metrics, err := h.metricUC.GetMetrics(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.MetricsResponse{
		TotalContainers:   int32(metrics.TotalContainers),
		RunningContainers: int32(metrics.RunningContainers),
		HealthPercentage:  int32(metrics.HealthPercentage),
		CpuUsage:          metrics.CPUUsage,
		MemoryUsage:       metrics.MemoryUsage,
	}, nil
}

func (h *grpcHandler) GetTools(ctx context.Context, req *pb.ToolsRequest) (*pb.ToolsResponse, error) {
	tools, err := h.toolUC.GetTools(ctx)
	if err != nil {
		return nil, err
	}

	var pbTools []*pb.Tool
	for _, t := range tools {
		pbTools = append(pbTools, &pb.Tool{
			Name:        t.Name,
			Subdomain:   t.Subdomain,
			Icon:        t.Icon,
			Description: t.Description,
			Protocols:   t.Protocols,
			Port:        t.Port,
			Uri:         t.URI,
			Category:    t.Category,
			Credentials: &pb.Credentials{
				Username: t.Credentials.Username,
				Password: t.Credentials.Password,
			},
			Uptime:  t.Uptime,
			Stopped: t.Stopped,
		})
	}

	return &pb.ToolsResponse{Tools: pbTools}, nil
}

func (h *grpcHandler) ControlTool(ctx context.Context, req *pb.ControlToolRequest) (*pb.ControlToolResponse, error) {
	var err error
	switch req.Action {
	case "start":
		err = h.toolUC.StartTool(ctx, req.Subdomain)
	case "stop":
		err = h.toolUC.StopTool(ctx, req.Subdomain)
	case "restart":
		err = h.toolUC.RestartTool(ctx, req.Subdomain)
	default:
		return nil, fmt.Errorf("invalid action: %s", req.Action)
	}

	if err != nil {
		return nil, err
	}

	return &pb.ControlToolResponse{Status: "success"}, nil
}

func (h *grpcHandler) StreamState(req *pb.StreamStateRequest, stream pb.AttackForgeService_StreamStateServer) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			metrics, err := h.GetMetrics(ctx, &pb.MetricsRequest{})
			if err != nil {
				cancel()
				return err
			}

			toolsResp, err := h.GetTools(ctx, &pb.ToolsRequest{})
			cancel()
			if err != nil {
				return err
			}

			err = stream.Send(&pb.StreamStateResponse{
				Metrics: metrics,
				Tools:   toolsResp.Tools,
			})
			if err != nil {
				return err
			}
		}
	}
}
