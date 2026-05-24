package patient_status_batch_get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCompute = errors.NewGRPCError(codes.Internal, "failed to compute patient statuses")
)
