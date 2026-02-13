package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreateSensorPatient = errors.NewGRPCError(codes.Internal, "failed to create sensor patient")
	errInvalidRequest              = errors.NewGRPCError(codes.InvalidArgument, "invalid request: sensor_id and patient_id are required")
	errSensorNotFound              = errors.NewGRPCError(codes.NotFound, "sensor not found")
)
