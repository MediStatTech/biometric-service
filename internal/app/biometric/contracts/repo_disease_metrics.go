package contracts

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type DiseaseMetricsRepo interface {
	FindByDiseaseID(ctx context.Context, diseaseID string) ([]domain.DiseaseMetricProps, error)
	FindAll(ctx context.Context) ([]domain.DiseaseMetricProps, error)
	CreateMut(diseaseMetric *domain.DiseaseMetric) *postgres.Mutation
	UpdateMut(diseaseMetric *domain.DiseaseMetric) *postgres.Mutation
}
