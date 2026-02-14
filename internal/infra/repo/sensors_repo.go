package repo

import (
	"context"
	"database/sql"
	"time"

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

// ============================================================================
// Sensor Patient Metrics Repository
// ============================================================================

type SensorPatientMetricsRepository struct {
	queries *Queries
}

func NewSensorPatientMetricsRepository(db *sql.DB) *SensorPatientMetricsRepository {
	return &SensorPatientMetricsRepository{
		queries: New(db),
	}
}

func (r *SensorPatientMetricsRepository) FindBySensorAndPatient(ctx context.Context, sensorID, patientID string) ([]domain.SensorPatientMetricProps, error) {
	sid, err := uuid.Parse(sensorID)
	if err != nil {
		return nil, err
	}
	pid, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	metrics, err := r.queries.ListSensorPatientMetrics(ctx, ListSensorPatientMetricsParams{
		SensorID:  sid,
		PatientID: pid,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.SensorPatientMetricProps, 0, len(metrics))
	for _, metric := range metrics {
		result = append(result, toSensorPatientMetricProps(metric))
	}
	return result, nil
}

func (r *SensorPatientMetricsRepository) FindByTimeRange(ctx context.Context, sensorID, patientID string, startTime, endTime time.Time) ([]domain.SensorPatientMetricProps, error) {
	sid, err := uuid.Parse(sensorID)
	if err != nil {
		return nil, err
	}
	pid, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	metrics, err := r.queries.ListSensorPatientMetricsByTimeRange(ctx, ListSensorPatientMetricsByTimeRangeParams{
		SensorID:    sid,
		PatientID:   pid,
		CreatedAt:   startTime,
		CreatedAt_2: endTime,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.SensorPatientMetricProps, 0, len(metrics))
	for _, metric := range metrics {
		result = append(result, toSensorPatientMetricProps(metric))
	}
	return result, nil
}

func (r *SensorPatientMetricsRepository) CreateMut(metric *domain.SensorPatientMetric) *postgres.Mutation {
	return postgres.NewMutation(
		CreateSensorPatientMetric,
		sensorPatientMetricToCreateParams(metric)...,
	)
}

func (r *SensorPatientMetricsRepository) DeleteBySensorAndPatientMut(sensorID, patientID string) *postgres.Mutation {
	sid, _ := uuid.Parse(sensorID)
	pid, _ := uuid.Parse(patientID)
	return postgres.NewMutation(
		DeleteSensorPatientMetrics,
		sid,
		pid,
	)
}

func (r *SensorPatientMetricsRepository) CreateBatchMut(metrics []*domain.SensorPatientMetric) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(metrics))
	for _, metric := range metrics {
		mutations = append(mutations, r.CreateMut(metric))
	}
	return mutations
}
