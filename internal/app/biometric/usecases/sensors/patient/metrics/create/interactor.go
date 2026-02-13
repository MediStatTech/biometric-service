package create

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
)

type Interactor struct {
	sensorsRepo              contracts.SensorsRepo
	sensorPatientsRepo       contracts.SensorPatientsRepo
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo
	committer                contracts.Committer
	logger                   contracts.Logger
}

func New(
	sensorsRepo contracts.SensorsRepo,
	sensorPatientsRepo contracts.SensorPatientsRepo,
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorsRepo:              sensorsRepo,
		sensorPatientsRepo:       sensorPatientsRepo,
		sensorPatientMetricsRepo: sensorPatientMetricsRepo,
		committer:                committer,
		logger:                   logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	//todo

	return nil, nil
}
