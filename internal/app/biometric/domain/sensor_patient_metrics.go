package domain

import "time"

type SensorPatientMetricProps struct {
	SensorID  string
	PatientID string
	MetricID  string
	Value     float64
	CreatedAt time.Time
}

type SensorPatientMetric struct {
	sensorID  string
	patientID string
	metricID  string
	value     float64
	createdAt time.Time
}

func NewSensorPatientMetric(
	sensorID string,
	patientID string,
	metricID string,
	value float64,
	createdAt time.Time,
) *SensorPatientMetric {
	return &SensorPatientMetric{
		sensorID:  sensorID,
		patientID: patientID,
		metricID:  metricID,
		value:     value,
		createdAt: createdAt,
	}
}

func ReconstituteSensorPatientMetric(p SensorPatientMetricProps) *SensorPatientMetric {
	return &SensorPatientMetric{
		sensorID:  p.SensorID,
		patientID: p.PatientID,
		metricID:  p.MetricID,
		value:     p.Value,
		createdAt: p.CreatedAt,
	}
}

func (spm *SensorPatientMetric) SensorID() string     { return spm.sensorID }
func (spm *SensorPatientMetric) PatientID() string    { return spm.patientID }
func (spm *SensorPatientMetric) MetricID() string     { return spm.metricID }
func (spm *SensorPatientMetric) Value() float64       { return spm.value }
func (spm *SensorPatientMetric) CreatedAt() time.Time { return spm.createdAt }

func (spm *SensorPatientMetric) SetValue(value float64) *SensorPatientMetric {
	spm.value = value
	return spm
}
