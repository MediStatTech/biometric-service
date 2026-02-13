package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

// ============================================================================
// Diseases Repository
// ============================================================================

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

	return toDiseaseProps(disease), nil
}

func (r *DiseasesRepository) FindByCode(ctx context.Context, code string) (domain.DiseaseProps, error) {
	disease, err := r.queries.GetDiseaseByCode(ctx, code)
	if err != nil {
		return domain.DiseaseProps{}, err
	}

	return toDiseaseProps(disease), nil
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

func (r *DiseasesRepository) CreateBatchMut(diseases []*domain.Disease) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(diseases))
	for _, disease := range diseases {
		mutations = append(mutations, r.CreateMut(disease))
	}
	return mutations
}

// ============================================================================
// Disease Sensors Repository
// ============================================================================

type DiseaseSensorsRepository struct {
	queries *Queries
}

func NewDiseaseSensorsRepository(db *sql.DB) *DiseaseSensorsRepository {
	return &DiseaseSensorsRepository{
		queries: New(db),
	}
}

func (r *DiseaseSensorsRepository) FindByDiseaseID(ctx context.Context, diseaseID string) ([]domain.DiseaseSensorProps, error) {
	id, err := uuid.Parse(diseaseID)
	if err != nil {
		return nil, err
	}

	diseaseSensors, err := r.queries.ListDiseaseSensorsByDisease(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.DiseaseSensorProps, 0, len(diseaseSensors))
	for _, ds := range diseaseSensors {
		result = append(result, toDiseaseSensorProps(ds))
	}

	return result, nil
}

func (r *DiseaseSensorsRepository) FindBySensorID(ctx context.Context, sensorID string) ([]domain.DiseaseSensorProps, error) {
	id, err := uuid.Parse(sensorID)
	if err != nil {
		return nil, err
	}

	diseaseSensors, err := r.queries.ListDiseaseSensorsBySensor(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.DiseaseSensorProps, 0, len(diseaseSensors))
	for _, ds := range diseaseSensors {
		result = append(result, toDiseaseSensorProps(ds))
	}

	return result, nil
}

func (r *DiseaseSensorsRepository) FindByDiseaseAndSensor(ctx context.Context, diseaseID, sensorID string) (domain.DiseaseSensorProps, error) {
	did, err := uuid.Parse(diseaseID)
	if err != nil {
		return domain.DiseaseSensorProps{}, err
	}

	sid, err := uuid.Parse(sensorID)
	if err != nil {
		return domain.DiseaseSensorProps{}, err
	}

	ds, err := r.queries.GetDiseaseSensor(ctx, GetDiseaseSensorParams{
		DiseaseID: did,
		SensorID:  sid,
	})
	if err != nil {
		return domain.DiseaseSensorProps{}, err
	}

	return toDiseaseSensorProps(ds), nil
}

func (r *DiseaseSensorsRepository) CountByDisease(ctx context.Context, diseaseID string) (int64, error) {
	id, err := uuid.Parse(diseaseID)
	if err != nil {
		return 0, err
	}

	return r.queries.CountDiseaseSensorsByDisease(ctx, id)
}

func (r *DiseaseSensorsRepository) CountBySensor(ctx context.Context, sensorID string) (int64, error) {
	id, err := uuid.Parse(sensorID)
	if err != nil {
		return 0, err
	}

	return r.queries.CountDiseaseSensorsBySensor(ctx, id)
}

func (r *DiseaseSensorsRepository) CreateMut(diseaseSensor *domain.DiseaseSensor) *postgres.Mutation {
	return postgres.NewMutation(
		CreateDiseaseSensor,
		diseaseSensorToCreateParams(diseaseSensor)...,
	)
}

func (r *DiseaseSensorsRepository) UpdateMut(diseaseSensor *domain.DiseaseSensor) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateDiseaseSensor,
		diseaseSensorToUpdateParams(diseaseSensor)...,
	)
}

func (r *DiseaseSensorsRepository) DeleteMut(diseaseID, sensorID string) *postgres.Mutation {
	did, _ := uuid.Parse(diseaseID)
	sid, _ := uuid.Parse(sensorID)
	return postgres.NewMutation(
		DeleteDiseaseSensor,
		DeleteDiseaseSensorParams{
			DiseaseID: did,
			SensorID:  sid,
		},
	)
}

func (r *DiseaseSensorsRepository) CreateBatchMut(diseaseSensors []*domain.DiseaseSensor) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(diseaseSensors))
	for _, ds := range diseaseSensors {
		mutations = append(mutations, r.CreateMut(ds))
	}
	return mutations
}
