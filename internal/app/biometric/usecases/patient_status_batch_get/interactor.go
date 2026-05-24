package patient_status_batch_get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
)

const (
	StatusOK       = "ok"
	StatusWarning  = "warning"
	StatusCritical = "critical"

	tolerance = 0.2
)

type Interactor struct {
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo
	metricTypesRepo          contracts.MetricTypesRepo
	logger                   contracts.Logger
}

func New(
	sensorPatientMetricsRepo contracts.SensorPatientMetricsRepo,
	metricTypesRepo contracts.MetricTypesRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorPatientMetricsRepo: sensorPatientMetricsRepo,
		metricTypesRepo:          metricTypesRepo,
		logger:                   logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	statuses := make([]PatientStatus, 0, len(req.PatientIDs))
	if len(req.PatientIDs) == 0 {
		return &Response{Statuses: statuses}, nil
	}

	latest, err := it.sensorPatientMetricsRepo.FindLatestForPatients(ctx, req.PatientIDs)
	if err != nil {
		return nil, errFailedToCompute.SetInternal(err)
	}

	metricTypes, err := it.metricTypesRepo.FindAll(ctx)
	if err != nil {
		return nil, errFailedToCompute.SetInternal(err)
	}
	typesByID := make(map[string]domain.MetricTypeProps, len(metricTypes))
	for _, mt := range metricTypes {
		typesByID[mt.MetricTypeID] = mt
	}

	patientWorst := make(map[string]string)
	for _, m := range latest {
		mt, ok := typesByID[m.MetricID]
		if !ok {
			continue
		}
		s := evaluate(m.Value, mt.MinValue, mt.MaxValue)
		cur, exists := patientWorst[m.PatientID]
		if !exists || rank(s) > rank(cur) {
			patientWorst[m.PatientID] = s
		}
	}

	for _, pid := range req.PatientIDs {
		s, ok := patientWorst[pid]
		if !ok {
			s = StatusOK
		}
		statuses = append(statuses, PatientStatus{
			PatientID: pid,
			Status:    s,
		})
	}

	return &Response{Statuses: statuses}, nil
}

func evaluate(value, minV, maxV float64) string {
	if maxV <= minV {
		return StatusOK
	}
	band := (maxV - minV) * tolerance
	if value >= minV && value <= maxV {
		return StatusOK
	}
	if value >= minV-band && value <= maxV+band {
		return StatusWarning
	}
	return StatusCritical
}

func rank(s string) int {
	switch s {
	case StatusCritical:
		return 3
	case StatusWarning:
		return 2
	case StatusOK:
		return 1
	}
	return 0
}
