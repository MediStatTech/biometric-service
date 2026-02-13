package disease_sensor_create

import (
	"context"
	"database/sql"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
)

type Interactor struct {
	diseasesRepo       contracts.DiseasesRepo
	sensorsRepo        contracts.SensorsRepo
	diseaseSensorsRepo contracts.DiseaseSensorsRepo
	committer          contracts.Committer
	logger             contracts.Logger
}

func New(
	diseasesRepo contracts.DiseasesRepo,
	sensorsRepo contracts.SensorsRepo,
	diseaseSensorsRepo contracts.DiseaseSensorsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseasesRepo:       diseasesRepo,
		sensorsRepo:        sensorsRepo,
		diseaseSensorsRepo: diseaseSensorsRepo,
		committer:          committer,
		logger:             logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.DiseaseID == "" || req.SensorID == "" {
		return nil, errInvalidRequest
	}

	_, err := it.diseasesRepo.FindByID(ctx, req.DiseaseID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errDiseaseNotFound
		}
		return nil, errFailedToCreateDiseaseSensor.SetInternal(err)
	}

	_, err = it.sensorsRepo.FindByID(ctx, req.SensorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errSensorNotFound
		}
		return nil, errFailedToCreateDiseaseSensor.SetInternal(err)
	}

	_, err = it.diseaseSensorsRepo.FindByDiseaseAndSensor(ctx, req.DiseaseID, req.SensorID)
	if err == nil {
		return nil, errDiseaseSensorExists
	}
	if err != sql.ErrNoRows {
		return nil, errFailedToCreateDiseaseSensor.SetInternal(err)
	}

	now := time.Now().UTC()
	diseaseSensor := domain.NewDiseaseSensor(
		req.DiseaseID,
		req.SensorID,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.diseaseSensorsRepo.CreateMut(diseaseSensor))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreateDiseaseSensor.SetInternal(err)
	}

	return &Response{
		DiseaseID: diseaseSensor.DiseaseID(),
		SensorID:  diseaseSensor.SensorID(),
	}, nil
}
