package get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetMetrics = errors.NewGRPCError(codes.Internal, "failed to get sensor patient metrics")
	errInvalidRequest     = errors.NewGRPCError(codes.InvalidArgument, "invalid request: sensor_id and patient_id are required")
)
