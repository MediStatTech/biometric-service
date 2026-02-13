package retrieve

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errInvalidRequest             = errors.NewGRPCError(codes.InvalidArgument, "invalid request: disease_id and sensor_id are required")
	errDiseaseSensorNotFound      = errors.NewGRPCError(codes.NotFound, "disease sensor not found")
	errFailedToGetDiseaseSensor   = errors.NewGRPCError(codes.Internal, "failed to retrieve disease sensor")
)
