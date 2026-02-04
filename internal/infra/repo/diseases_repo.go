package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

type DiseasesRepository struct {
	queries *Queries
}

func NewDiseasesRepository(db *sql.DB) *DiseasesRepository {
	return &DiseasesRepository{
		queries: New(db),
	}
}

func (r *DiseasesRepository) FindAll(ctx context.Context) ([]domain.DiseaseProps, error) {
	diseases, err := r.queries.ListDiseases(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.DiseaseProps, 0, len(diseases))
	for _, disease := range diseases {
		result = append(result, toDiseaseProps(disease))
	}

	return result, nil
}

func (r *DiseasesRepository) FindByID(ctx context.Context, diseaseID string) (domain.DiseaseProps, error) {
	id, err := uuid.Parse(diseaseID)
	if err != nil {
		return domain.DiseaseProps{}, err
	}

	disease, err := r.queries.GetDisease(ctx, id)
	if err != nil {
		return domain.DiseaseProps{}, err
	}

	return toDiseasePropsFromGetRow(disease), nil
}

func (r *DiseasesRepository) CreateMut(disease *domain.Disease) *postgres.Mutation {
	return postgres.NewMutation(
		CreateDisease,
		diseaseToCreateParams(disease)...,
	)
}

func (r *DiseasesRepository) UpdateMut(disease *domain.Disease) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateDisease,
		diseaseToUpdateParams(disease)...,
	)
}
