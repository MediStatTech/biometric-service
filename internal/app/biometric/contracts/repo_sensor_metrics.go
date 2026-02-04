package contracts

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type SensorMetricsRepo interface {
	FindBySensorID(ctx context.Context, sensorID string) ([]domain.SensorMetricProps, error)
	FindAll(ctx context.Context) ([]domain.SensorMetricProps, error)
	CreateMut(sensorMetric *domain.SensorMetric) *postgres.Mutation
	UpdateMut(sensorMetric *domain.SensorMetric) *postgres.Mutation
}
