package get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetSensorPatients = errors.NewGRPCError(codes.Internal, "failed to get sensor patients")
	errInvalidStatus             = errors.NewGRPCError(codes.InvalidArgument, "invalid status: must be 'active' or 'inactive'")
)
