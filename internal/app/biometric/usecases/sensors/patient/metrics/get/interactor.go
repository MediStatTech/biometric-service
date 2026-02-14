package get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

type Interactor struct {
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo
	logger                   contracts.Logger
}

func New(
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorPatientMetricsRepo: sensorPatientMetricsRepo,
		logger:                   logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.SensorID == "" || req.PatientID == "" {
		return nil, errInvalidRequest
	}

	metrics, err := it.sensorPatientMetricsRepo.FindBySensorAndPatient(ctx, req.SensorID, req.PatientID)
	if err != nil {
		return nil, errFailedToGetMetrics.SetInternal(err)
	}

	return &Response{
		Metrics: metrics,
	}, nil
}
