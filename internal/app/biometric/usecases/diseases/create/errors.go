package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetDiseaseMetrics = errors.NewGRPCError(codes.Internal, "failed to get disease metrics")
	errNoDiseaseMetricsFound     = errors.NewGRPCError(codes.NotFound, "no disease metrics found for disease")
	errFailedToCreateSensors     = errors.NewGRPCError(codes.Internal, "failed to create sensors")
)
