package disease_create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreateDisease = errors.NewGRPCError(codes.Internal, "failed to create disease")
	errInvalidRequest        = errors.NewGRPCError(codes.InvalidArgument, "invalid request: name and code are required")
	errDiseaseCodeExists     = errors.NewGRPCError(codes.AlreadyExists, "disease with this code already exists")
)
