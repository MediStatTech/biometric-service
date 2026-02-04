package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

type SensorMetricsRepository struct {
	queries *Queries
}

func NewSensorMetricsRepository(db *sql.DB) *SensorMetricsRepository {
	return &SensorMetricsRepository{
		queries: New(db),
	}
}

func (r *SensorMetricsRepository) FindBySensorID(ctx context.Context, sensorID string) ([]domain.SensorMetricProps, error) {
	id, err := uuid.Parse(sensorID)
	if err != nil {
		return nil, err
	}

	metrics, err := r.queries.ListSensorMetricsBySensorID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.SensorMetricProps, 0, len(metrics))
	for _, metric := range metrics {
		result = append(result, toSensorMetricPropsFromListBySensorIDRow(metric))
	}

	return result, nil
}

func (r *SensorMetricsRepository) FindAll(ctx context.Context) ([]domain.SensorMetricProps, error) {
	metrics, err := r.queries.ListSensorMetrics(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.SensorMetricProps, 0, len(metrics))
	for _, metric := range metrics {
		result = append(result, toSensorMetricProps(metric))
	}

	return result, nil
}

func (r *SensorMetricsRepository) CreateMut(sensorMetric *domain.SensorMetric) *postgres.Mutation {
	return postgres.NewMutation(
		CreateSensorMetric,
		sensorMetricToCreateParams(sensorMetric)...,
	)
}

func (r *SensorMetricsRepository) UpdateMut(sensorMetric *domain.SensorMetric) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateSensorMetric,
		sensorMetricToUpdateParams(sensorMetric)...,
	)
}
