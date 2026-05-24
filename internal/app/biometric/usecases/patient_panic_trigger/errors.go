package patient_panic_trigger

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errInvalidRequest = errors.NewGRPCError(codes.InvalidArgument, "invalid request: patient_id is required")
)
