package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreateMetricType = errors.NewGRPCError(codes.Internal, "failed to create metric type")
	errInvalidRequest           = errors.NewGRPCError(codes.InvalidArgument, "invalid request: sensor_id, code and name are required")
	errInvalidRange             = errors.NewGRPCError(codes.InvalidArgument, "invalid range: min_value must be less than max_value")
	errMetricTypeCodeExists     = errors.NewGRPCError(codes.AlreadyExists, "metric type with this code already exists for the sensor")
)
