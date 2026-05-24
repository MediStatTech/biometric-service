package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

type MetricTypesRepository struct {
	queries *Queries
}

func NewMetricTypesRepository(db *sql.DB) *MetricTypesRepository {
	return &MetricTypesRepository{
		queries: New(db),
	}
}

func (r *MetricTypesRepository) FindAll(ctx context.Context) ([]domain.MetricTypeProps, error) {
	rows, err := r.queries.ListMetricTypes(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.MetricTypeProps, 0, len(rows))
	for _, mt := range rows {
		result = append(result, toMetricTypeProps(mt))
	}
	return result, nil
}

func (r *MetricTypesRepository) FindByID(ctx context.Context, metricTypeID string) (domain.MetricTypeProps, error) {
	id, err := uuid.Parse(metricTypeID)
	if err != nil {
		return domain.MetricTypeProps{}, err
	}

	mt, err := r.queries.GetMetricType(ctx, id)
	if err != nil {
		return domain.MetricTypeProps{}, err
	}
	return toMetricTypeProps(mt), nil
}

func (r *MetricTypesRepository) FindBySensorID(ctx context.Context, sensorID string) ([]domain.MetricTypeProps, error) {
	id, err := uuid.Parse(sensorID)
	if err != nil {
		return nil, err
	}

	rows, err := r.queries.ListMetricTypesBySensor(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.MetricTypeProps, 0, len(rows))
	for _, mt := range rows {
		result = append(result, toMetricTypeProps(mt))
	}
	return result, nil
}

func (r *MetricTypesRepository) FindBySensorAndCode(ctx context.Context, sensorID, code string) (domain.MetricTypeProps, error) {
	sid, err := uuid.Parse(sensorID)
	if err != nil {
		return domain.MetricTypeProps{}, err
	}

	mt, err := r.queries.GetMetricTypeBySensorAndCode(ctx, GetMetricTypeBySensorAndCodeParams{
		SensorID: sid,
		Code:     code,
	})
	if err != nil {
		return domain.MetricTypeProps{}, err
	}
	return toMetricTypeProps(mt), nil
}

func (r *MetricTypesRepository) CountBySensor(ctx context.Context, sensorID string) (int64, error) {
	id, err := uuid.Parse(sensorID)
	if err != nil {
		return 0, err
	}
	return r.queries.CountMetricTypesBySensor(ctx, id)
}

func (r *MetricTypesRepository) CreateMut(metricType *domain.MetricType) *postgres.Mutation {
	return postgres.NewMutation(
		CreateMetricType,
		metricTypeToCreateParams(metricType)...,
	)
}

func (r *MetricTypesRepository) UpdateMut(metricType *domain.MetricType) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateMetricType,
		metricTypeToUpdateParams(metricType)...,
	)
}

func (r *MetricTypesRepository) DeleteMut(metricTypeID string) *postgres.Mutation {
	id, _ := uuid.Parse(metricTypeID)
	return postgres.NewMutation(DeleteMetricType, id)
}

func (r *MetricTypesRepository) CreateBatchMut(metricTypes []*domain.MetricType) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(metricTypes))
	for _, mt := range metricTypes {
		mutations = append(mutations, r.CreateMut(mt))
	}
	return mutations
}
