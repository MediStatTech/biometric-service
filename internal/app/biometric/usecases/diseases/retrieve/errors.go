package retrieve

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errInvalidRequest        = errors.NewGRPCError(codes.InvalidArgument, "invalid request: disease_id is required")
	errDiseaseNotFound       = errors.NewGRPCError(codes.NotFound, "disease not found")
	errFailedToGetDisease    = errors.NewGRPCError(codes.Internal, "failed to retrieve disease")
)
