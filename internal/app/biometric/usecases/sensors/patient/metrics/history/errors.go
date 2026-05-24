package history

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetMetrics = errors.NewGRPCError(codes.Internal, "failed to get metric history")
	errInvalidRequest     = errors.NewGRPCError(codes.InvalidArgument, "invalid request: sensor_id and patient_id are required")
	errInvalidTimeRange   = errors.NewGRPCError(codes.InvalidArgument, "invalid time range: start_time must be before end_time")
)
