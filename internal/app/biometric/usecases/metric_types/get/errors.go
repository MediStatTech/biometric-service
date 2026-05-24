package get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetMetricTypes = errors.NewGRPCError(codes.Internal, "failed to get metric types")
)
