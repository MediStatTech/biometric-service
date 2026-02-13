package get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetDiseaseSensors = errors.NewGRPCError(codes.Internal, "failed to get disease sensors")
)
