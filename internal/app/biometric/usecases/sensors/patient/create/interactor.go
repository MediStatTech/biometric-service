package create

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
)

type Interactor struct {
	sensorsRepo        contracts.SensorsRepo
	sensorPatientsRepo contracts.SensorPatientsRepo
	committer          contracts.Committer
	logger             contracts.Logger
}

func New(
	sensorsRepo contracts.SensorsRepo,
	sensorPatientsRepo contracts.SensorPatientsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorsRepo:        sensorsRepo,
		sensorPatientsRepo: sensorPatientsRepo,
		committer:          committer,
		logger:             logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	now := time.Now().UTC()
	sensorPatient := domain.NewSensorPatient(
		req.SensorID,
		req.PatientID,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.sensorPatientsRepo.CreateMut(sensorPatient))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreateSensorPatient.SetInternal(err)
	}

	return &Response{
		SensorID:  sensorPatient.SensorID(),
		PatientID: sensorPatient.PatientID(),
		Status:    sensorPatient.Status().String(),
	}, nil
}
