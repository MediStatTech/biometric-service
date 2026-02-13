package disease_get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetDiseases = errors.NewGRPCError(codes.Internal, "failed to get diseases")
)
