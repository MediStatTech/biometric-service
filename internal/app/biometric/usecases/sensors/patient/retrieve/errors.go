package retrieve

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetSensorPatient = errors.NewGRPCError(codes.Internal, "failed to get sensor patient")
	errInvalidRequest           = errors.NewGRPCError(codes.InvalidArgument, "invalid request: sensor_id and patient_id are required")
	errSensorPatientNotFound    = errors.NewGRPCError(codes.NotFound, "sensor patient not found")
)
