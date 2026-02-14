package repo

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/google/uuid"
)

// ============================================================================
// Diseases Mappers
// ============================================================================

func toDiseaseProps(disease Disease) domain.DiseaseProps {
	return domain.DiseaseProps{
		DiseaseID: disease.DiseaseID.String(),
		Name:      disease.Name,
		Code:      disease.Code,
		CreatedAt: disease.CreatedAt,
		UpdatedAt: disease.UpdatedAt,
	}
}

func diseaseToCreateParams(disease *domain.Disease) []any {
	id, _ := uuid.Parse(disease.DiseaseID())
	return []any{
		id,
		disease.Name(),
		disease.Code(),
		disease.CreatedAt(),
		disease.UpdatedAt(),
	}
}

func diseaseToUpdateParams(disease *domain.Disease) []any {
	id, _ := uuid.Parse(disease.DiseaseID())
	return []any{
		id,
		disease.Name(),
		disease.Code(),
		disease.UpdatedAt(),
	}
}

// ============================================================================
// Disease Sensors Mappers
// ============================================================================

func toDiseaseSensorProps(ds DiseaseSensor) domain.DiseaseSensorProps {
	return domain.DiseaseSensorProps{
		DiseaseID: ds.DiseaseID.String(),
		SensorID:  ds.SensorID.String(),
		CreatedAt: ds.CreatedAt,
		UpdatedAt: ds.UpdatedAt,
	}
}

func diseaseSensorToCreateParams(ds *domain.DiseaseSensor) []any {
	diseaseID, _ := uuid.Parse(ds.DiseaseID())
	sensorID, _ := uuid.Parse(ds.SensorID())
	return []any{
		diseaseID,
		sensorID,
		ds.CreatedAt(),
		ds.UpdatedAt(),
	}
}

func diseaseSensorToUpdateParams(ds *domain.DiseaseSensor) []any {
	diseaseID, _ := uuid.Parse(ds.DiseaseID())
	sensorID, _ := uuid.Parse(ds.SensorID())
	return []any{
		diseaseID,
		sensorID,
		ds.UpdatedAt(),
	}
}

// ============================================================================
// Sensors Mappers
// ============================================================================

func toSensorProps(sensor Sensor) domain.SensorProps {
	return domain.SensorProps{
		SensorID:  sensor.SensorID.String(),
		Name:      sensor.Name,
		Code:      sensor.Code,
		Symbol:    sensor.Symbol,
		CreatedAt: sensor.CreatedAt,
		UpdatedAt: sensor.UpdatedAt,
	}
}

func sensorToCreateParams(sensor *domain.Sensor) []any {
	id, _ := uuid.Parse(sensor.SensorID())
	return []any{
		id,
		sensor.Name(),
		sensor.Code(),
		sensor.Symbol(),
		sensor.CreatedAt(),
		sensor.UpdatedAt(),
	}
}

func sensorToUpdateParams(sensor *domain.Sensor) []any {
	id, _ := uuid.Parse(sensor.SensorID())
	return []any{
		id,
		sensor.Name(),
		sensor.Code(),
		sensor.Symbol(),
		sensor.UpdatedAt(),
	}
}

// ============================================================================
// Sensor Patients Mappers
// ============================================================================

func toSensorPatientProps(sp SensorPatient) domain.SensorPatientProps {
	return domain.SensorPatientProps{
		SensorID:  sp.SensorID.String(),
		PatientID: sp.PatientID.String(),
		Status:    sp.Status,
		CreatedAt: sp.CreatedAt,
		UpdatedAt: sp.UpdatedAt,
	}
}

func sensorPatientToCreateParams(sp *domain.SensorPatient) []any {
	sensorID, _ := uuid.Parse(sp.SensorID())
	patientID, _ := uuid.Parse(sp.PatientID())
	return []any{
		sensorID,
		patientID,
		sp.Status().String(),
		sp.CreatedAt(),
		sp.UpdatedAt(),
	}
}

func sensorPatientToUpdateParams(sp *domain.SensorPatient) []any {
	sensorID, _ := uuid.Parse(sp.SensorID())
	patientID, _ := uuid.Parse(sp.PatientID())
	return []any{
		sensorID,
		patientID,
		sp.Status().String(),
		sp.UpdatedAt(),
	}
}

// ============================================================================
// Sensor Patient Metrics Mappers
// ============================================================================

func toSensorPatientMetricProps(spm SensorPatientMetric) domain.SensorPatientMetricProps {
	return domain.SensorPatientMetricProps{
		SensorID:  spm.SensorID.String(),
		PatientID: spm.PatientID.String(),
		MetricID:  spm.MetricID.String(),
		Value:     spm.Value,
		Symbol:    spm.Symbol,
		CreatedAt: spm.CreatedAt,
	}
}

func sensorPatientMetricToCreateParams(spm *domain.SensorPatientMetric) []any {
	sensorID, _ := uuid.Parse(spm.SensorID())
	patientID, _ := uuid.Parse(spm.PatientID())
	metricID, _ := uuid.Parse(spm.MetricID())
	return []any{
		sensorID,
		patientID,
		metricID,
		spm.Value(),
		spm.Symbol(),
		spm.CreatedAt(),
	}
}
