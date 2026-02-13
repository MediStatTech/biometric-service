package retrieve

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errInvalidRequest     = errors.NewGRPCError(codes.InvalidArgument, "invalid request: sensor_id is required")
	errSensorNotFound     = errors.NewGRPCError(codes.NotFound, "sensor not found")
	errFailedToGetSensor  = errors.NewGRPCError(codes.Internal, "failed to retrieve sensor")
)
