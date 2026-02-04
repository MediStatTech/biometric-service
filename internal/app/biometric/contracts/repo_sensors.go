package contracts

import (
	"context"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type SensorsRepo interface {
	FindAll(ctx context.Context) ([]domain.SensorProps, error)
	FindByID(ctx context.Context, sensorID string) (domain.SensorProps, error)
	CreateMut(sensor *domain.Sensor) *postgres.Mutation
	UpdateMut(sensor *domain.Sensor) *postgres.Mutation
	CreateBatchMut(sensors []*domain.Sensor) []*postgres.Mutation
}
