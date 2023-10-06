package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/kioku-project/kioku/services/srs/proto"
)

type Health struct{}

func (h *Health) Check(_ context.Context, _ *pb.HealthCheckRequest, rsp *pb.HealthCheckResponse) error {
	rsp.Status = pb.HealthCheckResponse_SERVING
	return nil
}

func (h *Health) Watch(_ context.Context, _ *pb.HealthCheckRequest, _ pb.Health_WatchStream) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
