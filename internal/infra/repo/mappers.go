package repo

import (
	"database/sql"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	"github.com/google/uuid"
)

// ============================================================================
// Diseases Mappers
// ============================================================================

func toDiseaseProps(disease ListDiseasesRow) domain.DiseaseProps {
	props := domain.DiseaseProps{
		DiseaseID: disease.DiseaseID.String(),
		Name:      disease.Name,
		CreatedAt: disease.CreatedAt,
	}

	if disease.UpdatedAt.Valid {
		props.UpdatedAt = &disease.UpdatedAt.Time
	}

	return props
}

func toDiseasePropsFromGetRow(disease GetDiseaseRow) domain.DiseaseProps {
	props := domain.DiseaseProps{
		DiseaseID: disease.DiseaseID.String(),
		Name:      disease.Name,
		CreatedAt: disease.CreatedAt,
	}

	if disease.UpdatedAt.Valid {
		props.UpdatedAt = &disease.UpdatedAt.Time
	}

	return props
}

func diseaseToCreateParams(disease *domain.Disease) []any {
	id, _ := uuid.Parse(disease.DiseaseID())
	return []any{
		id,
		disease.Name(),
		disease.CreatedAt(),
	}
}

func diseaseToUpdateParams(disease *domain.Disease) []any {
	id, _ := uuid.Parse(disease.DiseaseID())
	var updatedAt sql.NullTime
	if disease.UpdatedAt() != nil {
		updatedAt = sql.NullTime{
			Time:  *disease.UpdatedAt(),
			Valid: true,
		}
	}

	return []any{
		id,
		disease.Name(),
		updatedAt,
	}
}

// ============================================================================
// DiseaseMetrics Mappers
// ============================================================================

func toDiseaseMetricProps(metric ListDiseaseMetricsRow) domain.DiseaseMetricProps {
	props := domain.DiseaseMetricProps{
		DiseaseID: metric.DiseaseID.String(),
		MetricID:  metric.MetricID.String(),
		Name:      metric.Name,
		EnumName:  metric.EnumName,
		CreatedAt: metric.CreatedAt,
	}

	if metric.UpdatedAt.Valid {
		props.UpdatedAt = &metric.UpdatedAt.Time
	}

	return props
}

func toDiseaseMetricPropsFromListByDiseaseIDRow(metric ListDiseaseMetricsByDiseaseIDRow) domain.DiseaseMetricProps {
	props := domain.DiseaseMetricProps{
		DiseaseID: metric.DiseaseID.String(),
		MetricID:  metric.MetricID.String(),
		Name:      metric.Name,
		EnumName:  metric.EnumName,
		CreatedAt: metric.CreatedAt,
	}

	if metric.UpdatedAt.Valid {
		props.UpdatedAt = &metric.UpdatedAt.Time
	}

	return props
}

func diseaseMetricToCreateParams(metric *domain.DiseaseMetric) []any {
	diseaseID, _ := uuid.Parse(metric.DiseaseID())
	metricID, _ := uuid.Parse(metric.MetricID())
	return []any{
		diseaseID,
		metricID,
		metric.Name(),
		metric.EnumName(),
		metric.CreatedAt(),
	}
}

func diseaseMetricToUpdateParams(metric *domain.DiseaseMetric) []any {
	diseaseID, _ := uuid.Parse(metric.DiseaseID())
	metricID, _ := uuid.Parse(metric.MetricID())
	var updatedAt sql.NullTime
	if metric.UpdatedAt() != nil {
		updatedAt = sql.NullTime{
			Time:  *metric.UpdatedAt(),
			Valid: true,
		}
	}

	return []any{
		diseaseID,
		metricID,
		metric.Name(),
		metric.EnumName(),
		updatedAt,
	}
}

// ============================================================================
// Sensors Mappers
// ============================================================================

func toSensorProps(sensor ListSensorsRow) domain.SensorProps {
	props := domain.SensorProps{
		SensorID:  sensor.SensorID.String(),
		Name:      sensor.Name,
		Status:    sensor.Status,
		EnumName:  sensor.EnumName,
		CreatedAt: sensor.CreatedAt,
	}

	if sensor.UpdatedAt.Valid {
		props.UpdatedAt = &sensor.UpdatedAt.Time
	}

	return props
}

func toSensorPropsFromGetRow(sensor GetSensorRow) domain.SensorProps {
	props := domain.SensorProps{
		SensorID:  sensor.SensorID.String(),
		Name:      sensor.Name,
		Status:    sensor.Status,
		EnumName:  sensor.EnumName,
		CreatedAt: sensor.CreatedAt,
	}

	if sensor.UpdatedAt.Valid {
		props.UpdatedAt = &sensor.UpdatedAt.Time
	}

	return props
}

func sensorToCreateParams(sensor *domain.Sensor) []any {
	id, _ := uuid.Parse(sensor.SensorID())
	return []any{
		id,
		sensor.Name(),
		sensor.Status().String(),
		sensor.EnumName(),
		sensor.CreatedAt(),
	}
}

func sensorToUpdateParams(sensor *domain.Sensor) []any {
	id, _ := uuid.Parse(sensor.SensorID())
	var updatedAt sql.NullTime
	if sensor.UpdatedAt() != nil {
		updatedAt = sql.NullTime{
			Time:  *sensor.UpdatedAt(),
			Valid: true,
		}
	}

	return []any{
		id,
		sensor.Name(),
		sensor.Status().String(),
		sensor.EnumName(),
		updatedAt,
	}
}

// ============================================================================
// SensorMetrics Mappers
// ============================================================================

func toSensorMetricProps(metric SensorMetric) domain.SensorMetricProps {
	props := domain.SensorMetricProps{
		SensorID:  metric.SensorID.String(),
		MetricID:  metric.MetricID.String(),
		Value:     metric.Value,
		CreatedAt: metric.CreatedAt,
	}

	return props
}

func toSensorMetricPropsFromListBySensorIDRow(metric SensorMetric) domain.SensorMetricProps {
	props := domain.SensorMetricProps{
		SensorID:  metric.SensorID.String(),
		MetricID:  metric.MetricID.String(),
		Value:     metric.Value,
		CreatedAt: metric.CreatedAt,
	}

	return props
}

func sensorMetricToCreateParams(metric *domain.SensorMetric) []any {
	sensorID, _ := uuid.Parse(metric.SensorID())
	metricID, _ := uuid.Parse(metric.MetricID())
	return []any{
		sensorID,
		metricID,
		metric.Value(),
		metric.CreatedAt(),
	}
}

func sensorMetricToUpdateParams(metric *domain.SensorMetric) []any {
	sensorID, _ := uuid.Parse(metric.SensorID())
	metricID, _ := uuid.Parse(metric.MetricID())
	return []any{
		sensorID,
		metricID,
		metric.Value(),
	}
}
