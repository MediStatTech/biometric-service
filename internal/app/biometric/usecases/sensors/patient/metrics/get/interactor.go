package get

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
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
	if req.SensorID == "" || req.PatientID == "" {
		return nil, errInvalidRequest
	}

	metrics, err := it.sensorPatientMetricsRepo.FindBySensorAndPatient(ctx, req.SensorID, req.PatientID)
	if err != nil {
		return nil, errFailedToGetMetrics.SetInternal(err)
	}
	if len(metrics) == 0 {
		return &Response{Measurements: nil}, nil
	}

	metricTypes, err := it.metricTypesRepo.FindBySensorID(ctx, req.SensorID)
	if err != nil {
		return nil, errFailedToGetMetrics.SetInternal(err)
	}

	typesByID := make(map[string]domain.MetricTypeProps, len(metricTypes))
	for _, mt := range metricTypes {
		typesByID[mt.MetricTypeID] = mt
	}

	groups := make(map[int64]*Measurement)
	order := make([]int64, 0)

	for _, m := range metrics {
		key := m.CreatedAt.UnixNano()
		g, ok := groups[key]
		if !ok {
			g = &Measurement{
				SensorID:  m.SensorID,
				PatientID: m.PatientID,
				CreatedAt: m.CreatedAt,
			}
			groups[key] = g
			order = append(order, key)
		}

		comp := Component{
			MetricTypeID: m.MetricID,
			Value:        m.Value,
			Symbol:       m.Symbol,
		}
		if mt, ok := typesByID[m.MetricID]; ok {
			comp.Code = mt.Code
			comp.Name = mt.Name
			if comp.Symbol == "" {
				comp.Symbol = mt.Symbol
			}
		}
		g.Components = append(g.Components, comp)
	}

	measurements := make([]Measurement, 0, len(order))
	for _, k := range order {
		measurements = append(measurements, *groups[k])
	}
	return &Response{Measurements: measurements}, nil
}
