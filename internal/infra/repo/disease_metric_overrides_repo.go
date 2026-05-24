package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

type DiseaseMetricOverridesRepository struct {
	queries *Queries
}

func NewDiseaseMetricOverridesRepository(db *sql.DB) *DiseaseMetricOverridesRepository {
	return &DiseaseMetricOverridesRepository{
		queries: New(db),
	}
}

func (r *DiseaseMetricOverridesRepository) FindAll(ctx context.Context) ([]domain.DiseaseMetricOverrideProps, error) {
	rows, err := r.queries.ListAllDiseaseMetricOverrides(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.DiseaseMetricOverrideProps, 0, len(rows))
	for _, o := range rows {
		result = append(result, toDiseaseMetricOverrideProps(o))
	}
	return result, nil
}

func (r *DiseaseMetricOverridesRepository) FindByDisease(ctx context.Context, diseaseID string) ([]domain.DiseaseMetricOverrideProps, error) {
	id, err := uuid.Parse(diseaseID)
	if err != nil {
		return nil, err
	}

	rows, err := r.queries.ListDiseaseMetricOverridesByDisease(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.DiseaseMetricOverrideProps, 0, len(rows))
	for _, o := range rows {
		result = append(result, toDiseaseMetricOverrideProps(o))
	}
	return result, nil
}

func (r *DiseaseMetricOverridesRepository) FindByMetricType(ctx context.Context, metricTypeID string) ([]domain.DiseaseMetricOverrideProps, error) {
	id, err := uuid.Parse(metricTypeID)
	if err != nil {
		return nil, err
	}

	rows, err := r.queries.ListDiseaseMetricOverridesByMetricType(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.DiseaseMetricOverrideProps, 0, len(rows))
	for _, o := range rows {
		result = append(result, toDiseaseMetricOverrideProps(o))
	}
	return result, nil
}

func (r *DiseaseMetricOverridesRepository) FindByDiseaseAndMetricType(ctx context.Context, diseaseID, metricTypeID string) (domain.DiseaseMetricOverrideProps, error) {
	did, err := uuid.Parse(diseaseID)
	if err != nil {
		return domain.DiseaseMetricOverrideProps{}, err
	}
	mid, err := uuid.Parse(metricTypeID)
	if err != nil {
		return domain.DiseaseMetricOverrideProps{}, err
	}

	o, err := r.queries.GetDiseaseMetricOverride(ctx, GetDiseaseMetricOverrideParams{
		DiseaseID:    did,
		MetricTypeID: mid,
	})
	if err != nil {
		return domain.DiseaseMetricOverrideProps{}, err
	}
	return toDiseaseMetricOverrideProps(o), nil
}

func (r *DiseaseMetricOverridesRepository) CreateMut(override *domain.DiseaseMetricOverride) *postgres.Mutation {
	return postgres.NewMutation(
		CreateDiseaseMetricOverride,
		diseaseMetricOverrideToCreateParams(override)...,
	)
}

func (r *DiseaseMetricOverridesRepository) UpdateMut(override *domain.DiseaseMetricOverride) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateDiseaseMetricOverride,
		diseaseMetricOverrideToUpdateParams(override)...,
	)
}

func (r *DiseaseMetricOverridesRepository) DeleteMut(diseaseID, metricTypeID string) *postgres.Mutation {
	did, _ := uuid.Parse(diseaseID)
	mid, _ := uuid.Parse(metricTypeID)
	return postgres.NewMutation(DeleteDiseaseMetricOverride, did, mid)
}

func (r *DiseaseMetricOverridesRepository) CreateBatchMut(overrides []*domain.DiseaseMetricOverride) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(overrides))
	for _, o := range overrides {
		mutations = append(mutations, r.CreateMut(o))
	}
	return mutations
}
