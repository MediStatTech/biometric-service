package contracts

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

// ============================================================================
// Sensors Repository
// ============================================================================

type SensorsRepo interface {
	FindAll(ctx context.Context) ([]domain.SensorProps, error)
	FindByID(ctx context.Context, sensorID string) (domain.SensorProps, error)
	FindByCode(ctx context.Context, code string) (domain.SensorProps, error)
	CreateMut(sensor *domain.Sensor) *postgres.Mutation
	UpdateMut(sensor *domain.Sensor) *postgres.Mutation
	CreateBatchMut(sensors []*domain.Sensor) []*postgres.Mutation
}

// ============================================================================
// Sensor Patients Repository
// ============================================================================

type SensorPatientsRepo interface {
	FindBySensorID(ctx context.Context, sensorID string) ([]domain.SensorPatientProps, error)
	FindByPatientID(ctx context.Context, patientID string) ([]domain.SensorPatientProps, error)
	FindByStatus(ctx context.Context, status domain.SensorPatientStatus) ([]domain.SensorPatientProps, error)
	FindBySensorAndPatient(ctx context.Context, sensorID, patientID string) (domain.SensorPatientProps, error)
	CountBySensor(ctx context.Context, sensorID string) (int64, error)
	CreateMut(sensorPatient *domain.SensorPatient) *postgres.Mutation
	UpdateMut(sensorPatient *domain.SensorPatient) *postgres.Mutation
	DeleteMut(sensorID, patientID string) *postgres.Mutation
	CreateBatchMut(sensorPatients []*domain.SensorPatient) []*postgres.Mutation
}

// ============================================================================
// Sensor Patient Metrics Repository
// ============================================================================

type SensorPatientMetricsRepo interface {
	FindBySensorAndPatient(ctx context.Context, sensorID, patientID string) ([]domain.SensorPatientMetricProps, error)
	FindByTimeRange(ctx context.Context, sensorID, patientID string, startTime, endTime time.Time) ([]domain.SensorPatientMetricProps, error)
	CreateMut(metric *domain.SensorPatientMetric) *postgres.Mutation
	DeleteBySensorAndPatientMut(sensorID, patientID string) *postgres.Mutation
	CreateBatchMut(metrics []*domain.SensorPatientMetric) []*postgres.Mutation
}
