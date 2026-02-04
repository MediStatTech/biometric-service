package contracts

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type DiseasesRepo interface {
	FindAll(ctx context.Context) ([]domain.DiseaseProps, error)
	FindByID(ctx context.Context, diseaseID string) (domain.DiseaseProps, error)
	CreateMut(disease *domain.Disease) *postgres.Mutation
	UpdateMut(disease *domain.Disease) *postgres.Mutation
}
