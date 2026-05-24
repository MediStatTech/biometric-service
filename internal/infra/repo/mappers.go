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

// ============================================================================
// Metric Types Mappers
// ============================================================================

func toMetricTypeProps(mt MetricType) domain.MetricTypeProps {
	return domain.MetricTypeProps{
		MetricTypeID: mt.MetricTypeID.String(),
		SensorID:     mt.SensorID.String(),
		Code:         mt.Code,
		Name:         mt.Name,
		Symbol:       mt.Symbol,
		MinValue:     mt.MinValue,
		MaxValue:     mt.MaxValue,
		CreatedAt:    mt.CreatedAt,
		UpdatedAt:    mt.UpdatedAt,
	}
}

func metricTypeToCreateParams(mt *domain.MetricType) []any {
	id, _ := uuid.Parse(mt.MetricTypeID())
	sensorID, _ := uuid.Parse(mt.SensorID())
	return []any{
		id,
		sensorID,
		mt.Code(),
		mt.Name(),
		mt.Symbol(),
		mt.MinValue(),
		mt.MaxValue(),
		mt.CreatedAt(),
		mt.UpdatedAt(),
	}
}

func metricTypeToUpdateParams(mt *domain.MetricType) []any {
	id, _ := uuid.Parse(mt.MetricTypeID())
	return []any{
		id,
		mt.Code(),
		mt.Name(),
		mt.Symbol(),
		mt.MinValue(),
		mt.MaxValue(),
		mt.UpdatedAt(),
	}
}

// ============================================================================
// Disease Metric Overrides Mappers
// ============================================================================

func toDiseaseMetricOverrideProps(o DiseaseMetricOverride) domain.DiseaseMetricOverrideProps {
	return domain.DiseaseMetricOverrideProps{
		DiseaseID:    o.DiseaseID.String(),
		MetricTypeID: o.MetricTypeID.String(),
		MinValue:     o.MinValue,
		MaxValue:     o.MaxValue,
		CreatedAt:    o.CreatedAt,
		UpdatedAt:    o.UpdatedAt,
	}
}

func diseaseMetricOverrideToCreateParams(o *domain.DiseaseMetricOverride) []any {
	did, _ := uuid.Parse(o.DiseaseID())
	mid, _ := uuid.Parse(o.MetricTypeID())
	return []any{
		did,
		mid,
		o.MinValue(),
		o.MaxValue(),
		o.CreatedAt(),
		o.UpdatedAt(),
	}
}

func diseaseMetricOverrideToUpdateParams(o *domain.DiseaseMetricOverride) []any {
	did, _ := uuid.Parse(o.DiseaseID())
	mid, _ := uuid.Parse(o.MetricTypeID())
	return []any{
		did,
		mid,
		o.MinValue(),
		o.MaxValue(),
		o.UpdatedAt(),
	}
}
