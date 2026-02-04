package domain

import (
	"time"

	"github.com/google/uuid"
)

type DiseaseMetricEnumName string

const (
	DiseaseMetricEnumNameTemperature   DiseaseMetricEnumName = "temperature"
	DiseaseMetricEnumNameHeartRate     DiseaseMetricEnumName = "heart_rate"
	DiseaseMetricEnumNameBloodPressure DiseaseMetricEnumName = "blood_pressure"
	DiseaseMetricEnumNameBloodOxygen   DiseaseMetricEnumName = "blood_oxygen"
)

func (d DiseaseMetricEnumName) String() string {
	return string(d)
}

func (d DiseaseMetricEnumName) Get() []string {
	return []string{
		DiseaseMetricEnumNameTemperature.String(),
		DiseaseMetricEnumNameHeartRate.String(),
		DiseaseMetricEnumNameBloodPressure.String(),
		DiseaseMetricEnumNameBloodOxygen.String(),
	}
}

type DiseaseMetricProps struct {
	DiseaseID string
	MetricID  string
	Name      string
	EnumName  string
	UpdatedAt *time.Time
	CreatedAt time.Time
}

type DiseaseMetric struct {
	diseaseID string
	metricID  string
	name      string
	enumName  string
	updatedAt *time.Time
	createdAt time.Time
}

func NewDiseaseMetric(
	diseaseID string,
	name string,
	createdAt time.Time,
	enumName string,
) *DiseaseMetric {
	return &DiseaseMetric{
		diseaseID: diseaseID,
		metricID:  uuid.NewString(),
		name:      name,
		enumName:  enumName,
		createdAt: createdAt,
	}
}

func ReconstituteDiseaseMetric(p DiseaseMetricProps) *DiseaseMetric {
	return &DiseaseMetric{
		diseaseID: p.DiseaseID,
		metricID:  p.MetricID,
		name:      p.Name,
		enumName:  p.EnumName,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (dm *DiseaseMetric) DiseaseID() string     { return dm.diseaseID }
func (dm *DiseaseMetric) MetricID() string      { return dm.metricID }
func (dm *DiseaseMetric) Name() string          { return dm.name }
func (dm *DiseaseMetric) EnumName() string      { return dm.enumName }
func (dm *DiseaseMetric) UpdatedAt() *time.Time { return dm.updatedAt }
func (dm *DiseaseMetric) CreatedAt() time.Time  { return dm.createdAt }

func (dm *DiseaseMetric) SetName(name string) *DiseaseMetric {
	dm.name = name
	return dm
}

func (dm *DiseaseMetric) SetEnumName(enumName string) *DiseaseMetric {
	dm.enumName = enumName
	return dm
}

func (dm *DiseaseMetric) SetUpdatedAt(updatedAt time.Time) *DiseaseMetric {
	dm.updatedAt = &updatedAt
	return dm
}
