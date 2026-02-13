package retrieve

import (
	"context"
	"database/sql"

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
	if req.SensorID == "" || req.PatientID == "" {
		return nil, errInvalidRequest
	}

	sensorPatient, err := it.sensorPatientsRepo.FindBySensorAndPatient(ctx, req.SensorID, req.PatientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errSensorPatientNotFound
		}
		return nil, errFailedToGetSensorPatient.SetInternal(err)
	}

	return &Response{
		SensorPatient: sensorPatient,
	}, nil
}
