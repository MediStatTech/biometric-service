package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreateSensor = errors.NewGRPCError(codes.Internal, "failed to create sensor")
	errInvalidRequest       = errors.NewGRPCError(codes.InvalidArgument, "invalid request: name and code are required")
)
