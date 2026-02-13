package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreateDiseaseSensor = errors.NewGRPCError(codes.Internal, "failed to create disease sensor")
	errInvalidRequest              = errors.NewGRPCError(codes.InvalidArgument, "invalid request: disease_id and sensor_id are required")
	errDiseaseNotFound             = errors.NewGRPCError(codes.NotFound, "disease not found")
	errSensorNotFound              = errors.NewGRPCError(codes.NotFound, "sensor not found")
	errDiseaseSensorExists         = errors.NewGRPCError(codes.AlreadyExists, "disease sensor relationship already exists")
)
