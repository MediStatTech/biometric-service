package contracts

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type MetricTypesRepo interface {
	FindAll(ctx context.Context) ([]domain.MetricTypeProps, error)
	FindByID(ctx context.Context, metricTypeID string) (domain.MetricTypeProps, error)
	FindBySensorID(ctx context.Context, sensorID string) ([]domain.MetricTypeProps, error)
	FindBySensorAndCode(ctx context.Context, sensorID, code string) (domain.MetricTypeProps, error)
	CountBySensor(ctx context.Context, sensorID string) (int64, error)
	CreateMut(metricType *domain.MetricType) *postgres.Mutation
	UpdateMut(metricType *domain.MetricType) *postgres.Mutation
	DeleteMut(metricTypeID string) *postgres.Mutation
	CreateBatchMut(metricTypes []*domain.MetricType) []*postgres.Mutation
}

type DiseaseMetricOverridesRepo interface {
	FindAll(ctx context.Context) ([]domain.DiseaseMetricOverrideProps, error)
	FindByDisease(ctx context.Context, diseaseID string) ([]domain.DiseaseMetricOverrideProps, error)
	FindByMetricType(ctx context.Context, metricTypeID string) ([]domain.DiseaseMetricOverrideProps, error)
	FindByDiseaseAndMetricType(ctx context.Context, diseaseID, metricTypeID string) (domain.DiseaseMetricOverrideProps, error)
	CreateMut(override *domain.DiseaseMetricOverride) *postgres.Mutation
	UpdateMut(override *domain.DiseaseMetricOverride) *postgres.Mutation
	DeleteMut(diseaseID, metricTypeID string) *postgres.Mutation
	CreateBatchMut(overrides []*domain.DiseaseMetricOverride) []*postgres.Mutation
}
