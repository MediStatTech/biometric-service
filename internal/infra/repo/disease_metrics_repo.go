package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

type DiseaseMetricsRepository struct {
	queries *Queries
}

func NewDiseaseMetricsRepository(db *sql.DB) *DiseaseMetricsRepository {
	return &DiseaseMetricsRepository{
		queries: New(db),
	}
}

func (r *DiseaseMetricsRepository) FindByDiseaseID(ctx context.Context, diseaseID string) ([]domain.DiseaseMetricProps, error) {
	id, err := uuid.Parse(diseaseID)
	if err != nil {
		return nil, err
	}

	metrics, err := r.queries.ListDiseaseMetricsByDiseaseID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.DiseaseMetricProps, 0, len(metrics))
	for _, metric := range metrics {
		result = append(result, toDiseaseMetricPropsFromListByDiseaseIDRow(metric))
	}

	return result, nil
}

func (r *DiseaseMetricsRepository) FindAll(ctx context.Context) ([]domain.DiseaseMetricProps, error) {
	metrics, err := r.queries.ListDiseaseMetrics(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.DiseaseMetricProps, 0, len(metrics))
	for _, metric := range metrics {
		result = append(result, toDiseaseMetricProps(metric))
	}

	return result, nil
}

func (r *DiseaseMetricsRepository) CreateMut(diseaseMetric *domain.DiseaseMetric) *postgres.Mutation {
	return postgres.NewMutation(
		CreateDiseaseMetric,
		diseaseMetricToCreateParams(diseaseMetric)...,
	)
}

func (r *DiseaseMetricsRepository) UpdateMut(diseaseMetric *domain.DiseaseMetric) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateDiseaseMetric,
		diseaseMetricToUpdateParams(diseaseMetric)...,
	)
}
