package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreateMetric   = errors.NewGRPCError(codes.Internal, "failed to create sensor patient metric")
	errInvalidRequest         = errors.NewGRPCError(codes.InvalidArgument, "invalid request: all fields are required")
	errSensorPatientNotFound  = errors.NewGRPCError(codes.NotFound, "sensor patient relationship not found")
	errSensorNotFoundByCode   = errors.NewGRPCError(codes.NotFound, "sensor not found by code")
)
