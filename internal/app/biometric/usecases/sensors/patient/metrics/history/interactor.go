package history

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/contracts"
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	metric_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/metrics/get"
)

const (
	defaultLimit int32 = 1000
	maxLimit     int32 = 5000
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
	if req.EndTime.IsZero() {
		req.EndTime = time.Now().UTC()
	}
	if req.StartTime.IsZero() || !req.StartTime.Before(req.EndTime) {
		return nil, errInvalidTimeRange
	}

	limit := req.Limit
	if limit <= 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset
	if offset < 0 {
		offset = 0
	}

	total, err := it.sensorPatientMetricsRepo.CountByTimeRange(ctx, req.SensorID, req.PatientID, req.StartTime, req.EndTime)
	if err != nil {
		return nil, errFailedToGetMetrics.SetInternal(err)
	}
	if total == 0 {
		return &Response{Measurements: nil, Total: 0}, nil
	}

	rows, err := it.sensorPatientMetricsRepo.FindByTimeRangePaged(ctx, req.SensorID, req.PatientID, req.StartTime, req.EndTime, limit, offset)
	if err != nil {
		return nil, errFailedToGetMetrics.SetInternal(err)
	}
	if len(rows) == 0 {
		return &Response{Measurements: nil, Total: total}, nil
	}

	metricTypes, err := it.metricTypesRepo.FindBySensorID(ctx, req.SensorID)
	if err != nil {
		return nil, errFailedToGetMetrics.SetInternal(err)
	}
	typesByID := make(map[string]domain.MetricTypeProps, len(metricTypes))
	for _, mt := range metricTypes {
		typesByID[mt.MetricTypeID] = mt
	}

	groups := make(map[int64]*metric_get.Measurement)
	order := make([]int64, 0)

	for _, m := range rows {
		key := m.CreatedAt.UnixNano()
		g, ok := groups[key]
		if !ok {
			g = &metric_get.Measurement{
				SensorID:  m.SensorID,
				PatientID: m.PatientID,
				CreatedAt: m.CreatedAt,
			}
			groups[key] = g
			order = append(order, key)
		}
		comp := metric_get.Component{
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

	measurements := make([]metric_get.Measurement, 0, len(order))
	for _, k := range order {
		measurements = append(measurements, *groups[k])
	}

	return &Response{
		Measurements: measurements,
		Total:        total,
	}, nil
}
