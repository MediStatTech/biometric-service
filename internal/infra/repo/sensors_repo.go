package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

// ============================================================================
// Sensors Repository
// ============================================================================

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

	return toSensorProps(sensor), nil
}

func (r *SensorsRepository) FindByCode(ctx context.Context, code string) (domain.SensorProps, error) {
	sensor, err := r.queries.GetSensorByCode(ctx, code)
	if err != nil {
		return domain.SensorProps{}, err
	}

	return toSensorProps(sensor), nil
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

// ============================================================================
// Sensor Patients Repository
// ============================================================================

type SensorPatientsRepository struct {
	queries *Queries
}

func NewSensorPatientsRepository(db *sql.DB) *SensorPatientsRepository {
	return &SensorPatientsRepository{
		queries: New(db),
	}
}

func (r *SensorPatientsRepository) FindBySensorID(ctx context.Context, sensorID string) ([]domain.SensorPatientProps, error) {
	id, err := uuid.Parse(sensorID)
	if err != nil {
		return nil, err
	}

	sensorPatients, err := r.queries.ListSensorPatientsBySensor(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.SensorPatientProps, 0, len(sensorPatients))
	for _, sp := range sensorPatients {
		result = append(result, toSensorPatientProps(sp))
	}

	return result, nil
}

func (r *SensorPatientsRepository) FindByPatientID(ctx context.Context, patientID string) ([]domain.SensorPatientProps, error) {
	id, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	sensorPatients, err := r.queries.ListSensorPatientsByPatient(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.SensorPatientProps, 0, len(sensorPatients))
	for _, sp := range sensorPatients {
		result = append(result, toSensorPatientProps(sp))
	}

	return result, nil
}

func (r *SensorPatientsRepository) FindByStatus(ctx context.Context, status domain.SensorPatientStatus) ([]domain.SensorPatientProps, error) {
	sensorPatients, err := r.queries.ListSensorPatientsByStatus(ctx, status.String())
	if err != nil {
		return nil, err
	}

	result := make([]domain.SensorPatientProps, 0, len(sensorPatients))
	for _, sp := range sensorPatients {
		result = append(result, toSensorPatientProps(sp))
	}

	return result, nil
}

func (r *SensorPatientsRepository) FindBySensorAndPatient(ctx context.Context, sensorID, patientID string) (domain.SensorPatientProps, error) {
	sid, err := uuid.Parse(sensorID)
	if err != nil {
		return domain.SensorPatientProps{}, err
	}

	pid, err := uuid.Parse(patientID)
	if err != nil {
		return domain.SensorPatientProps{}, err
	}

	sp, err := r.queries.GetSensorPatient(ctx, GetSensorPatientParams{
		SensorID:  sid,
		PatientID: pid,
	})
	if err != nil {
		return domain.SensorPatientProps{}, err
	}

	return toSensorPatientProps(sp), nil
}

func (r *SensorPatientsRepository) CountBySensor(ctx context.Context, sensorID string) (int64, error) {
	id, err := uuid.Parse(sensorID)
	if err != nil {
		return 0, err
	}

	count, err := r.queries.CountSensorPatientsBySensor(ctx, id)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *SensorPatientsRepository) CreateMut(sensorPatient *domain.SensorPatient) *postgres.Mutation {
	return postgres.NewMutation(
		CreateSensorPatient,
		sensorPatientToCreateParams(sensorPatient)...,
	)
}

func (r *SensorPatientsRepository) UpdateMut(sensorPatient *domain.SensorPatient) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateSensorPatient,
		sensorPatientToUpdateParams(sensorPatient)...,
	)
}

func (r *SensorPatientsRepository) DeleteMut(sensorID, patientID string) *postgres.Mutation {
	sid, _ := uuid.Parse(sensorID)
	pid, _ := uuid.Parse(patientID)
	return postgres.NewMutation(
		DeleteSensorPatient,
		sid,
		pid,
	)
}

func (r *SensorPatientsRepository) CreateBatchMut(sensorPatients []*domain.SensorPatient) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(sensorPatients))
	for _, sp := range sensorPatients {
		mutations = append(mutations, r.CreateMut(sp))
	}
	return mutations
}
