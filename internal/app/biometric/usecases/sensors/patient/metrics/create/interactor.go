package create

import (
	"context"
	"database/sql"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/biometric-service/pkg/commitplan"
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
	if req.SensorID == "" || req.PatientID == "" || req.MetricID == "" {
		return nil, errInvalidRequest
	}

	_, err := it.sensorPatientsRepo.FindBySensorAndPatient(ctx, req.SensorID, req.PatientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errSensorPatientNotFound
		}
		return nil, errFailedToCreateMetric.SetInternal(err)
	}

	//TODO

	now := time.Now().UTC()
	metric := domain.NewSensorPatientMetric(
		req.SensorID,
		req.PatientID,
		req.MetricID,
		req.Value,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.sensorPatientMetricsRepo.CreateMut(metric))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreateMetric.SetInternal(err)
	}

	return &Response{
		SensorID:  metric.SensorID(),
		PatientID: metric.PatientID(),
		MetricID:  metric.MetricID(),
		Value:     metric.Value(),
		CreatedAt: metric.CreatedAt().Format(time.RFC3339),
	}, nil
}
