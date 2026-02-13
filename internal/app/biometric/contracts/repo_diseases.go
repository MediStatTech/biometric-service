package contracts

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

// ============================================================================
// Diseases Repository
// ============================================================================

type DiseasesRepo interface {
	FindAll(ctx context.Context) ([]domain.DiseaseProps, error)
	FindByID(ctx context.Context, diseaseID string) (domain.DiseaseProps, error)
	FindByCode(ctx context.Context, code string) (domain.DiseaseProps, error)
	CreateMut(disease *domain.Disease) *postgres.Mutation
	UpdateMut(disease *domain.Disease) *postgres.Mutation
	CreateBatchMut(diseases []*domain.Disease) []*postgres.Mutation
}

// ============================================================================
// Disease Sensors Repository
// ============================================================================

type DiseaseSensorsRepo interface {
	FindByDiseaseID(ctx context.Context, diseaseID string) ([]domain.DiseaseSensorProps, error)
	FindBySensorID(ctx context.Context, sensorID string) ([]domain.DiseaseSensorProps, error)
	FindByDiseaseAndSensor(ctx context.Context, diseaseID, sensorID string) (domain.DiseaseSensorProps, error)
	CountByDisease(ctx context.Context, diseaseID string) (int64, error)
	CountBySensor(ctx context.Context, sensorID string) (int64, error)
	CreateMut(diseaseSensor *domain.DiseaseSensor) *postgres.Mutation
	UpdateMut(diseaseSensor *domain.DiseaseSensor) *postgres.Mutation
	DeleteMut(diseaseID, sensorID string) *postgres.Mutation
	CreateBatchMut(diseaseSensors []*domain.DiseaseSensor) []*postgres.Mutation
}
