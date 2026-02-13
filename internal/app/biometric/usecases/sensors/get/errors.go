package get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetSensors = errors.NewGRPCError(codes.Internal, "failed to get sensors")
)
