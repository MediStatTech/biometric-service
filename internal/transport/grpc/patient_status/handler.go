package patient_status

import (
	"context"

	patient_panic_trigger "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/patient_panic_trigger"
	patient_status_batch_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/patient_status_batch_get"
	pb_services "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) PatientStatusGetBatch(
	ctx context.Context,
	req *pb_services.PatientStatusGetBatchRequest,
) (*pb_services.PatientStatusGetBatchReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.PatientStatusGetBatch.Execute(ctx, patient_status_batch_get.Request{
		PatientIDs: req.GetPatientIds(),
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to compute patient statuses: %v", err)
		return nil, err
	}

	statuses := make([]*pb_services.PatientStatus, 0, len(resp.Statuses))
	for _, s := range resp.Statuses {
		statuses = append(statuses, &pb_services.PatientStatus{
			PatientId: s.PatientID,
			Status:    s.Status,
		})
	}

	return &pb_services.PatientStatusGetBatchReply{
		PatientStatuses: statuses,
	}, nil
}

func (h *Handler) PatientPanicTrigger(
	ctx context.Context,
	req *pb_services.PatientPanicTriggerRequest,
) (*pb_services.PatientPanicTriggerReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.commands.PatientPanicTrigger.Execute(ctx, patient_panic_trigger.Request{
		PatientID:       req.GetPatientId(),
		DurationSeconds: req.GetDurationSeconds(),
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to trigger patient panic: %v", err)
		return nil, err
	}

	return &pb_services.PatientPanicTriggerReply{
		PanicUntil: timestamppb.New(resp.PanicUntil),
	}, nil
}
