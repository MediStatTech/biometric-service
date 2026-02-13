package get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

type Interactor struct {
	sensorPatientsRepo contracts.SensorPatientsRepo
	logger             contracts.Logger
}

func New(
	sensorPatientsRepo contracts.SensorPatientsRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorPatientsRepo: sensorPatientsRepo,
		logger:             logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	sensorPatients, err := it.sensorPatientsRepo.FindBySensorID(ctx, req.SensorID)
	if err != nil {
		return nil, errFailedToGetSensorPatients.SetInternal(err)
	}

	return &Response{
		SensorPatients: sensorPatients,
	}, nil
}
