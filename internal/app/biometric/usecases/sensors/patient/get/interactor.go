package get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
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
	var sensorPatients []domain.SensorPatientProps
	var err error

	if req.Status != nil {
		status := domain.SensorPatientStatus(*req.Status)
		if status != domain.SensorPatientStatusActive && status != domain.SensorPatientStatusInactive {
			return nil, errInvalidStatus
		}
		sensorPatients, err = it.sensorPatientsRepo.FindByStatus(ctx, status)
	} else if req.SensorID != nil {
		sensorPatients, err = it.sensorPatientsRepo.FindBySensorID(ctx, *req.SensorID)
	} else if req.PatientID != nil {
		sensorPatients, err = it.sensorPatientsRepo.FindByPatientID(ctx, *req.PatientID)
	} else {
		sensorPatients, err = it.sensorPatientsRepo.FindByStatus(ctx, domain.SensorPatientStatusActive)
	}

	if err != nil {
		return nil, errFailedToGetSensorPatients.SetInternal(err)
	}

	return &Response{
		SensorPatients: sensorPatients,
	}, nil
}
