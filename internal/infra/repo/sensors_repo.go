package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

type SensorsRepository struct {
	queries *Queries
}

func NewSensorsRepository(db *sql.DB) *SensorsRepository {
	return &SensorsRepository{
		queries: New(db),
	}
}

func (r *SensorsRepository) FindAll(ctx context.Context) ([]domain.SensorProps, error) {
	sensors, err := r.queries.ListSensors(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.SensorProps, 0, len(sensors))
	for _, sensor := range sensors {
		result = append(result, toSensorProps(sensor))
	}

	return result, nil
}

func (r *SensorsRepository) FindByID(ctx context.Context, sensorID string) (domain.SensorProps, error) {
	id, err := uuid.Parse(sensorID)
	if err != nil {
		return domain.SensorProps{}, err
	}

	sensor, err := r.queries.GetSensor(ctx, id)
	if err != nil {
		return domain.SensorProps{}, err
	}

	return toSensorPropsFromGetRow(sensor), nil
}

func (r *SensorsRepository) CreateMut(sensor *domain.Sensor) *postgres.Mutation {
	return postgres.NewMutation(
		CreateSensor,
		sensorToCreateParams(sensor)...,
	)
}

func (r *SensorsRepository) UpdateMut(sensor *domain.Sensor) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateSensor,
		sensorToUpdateParams(sensor)...,
	)
}

func (r *SensorsRepository) CreateBatchMut(sensors []*domain.Sensor) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(sensors))
	for _, sensor := range sensors {
		mutations = append(mutations, r.CreateMut(sensor))
	}
	return mutations
}
