package domain

import "time"

type SensorPatientMetricProps struct {
	SensorID  string
	PatientID string
	MetricID  string
	Value     float64
	Symbol    string
	CreatedAt time.Time
}

type SensorPatientMetric struct {
	sensorID  string
	patientID string
	metricID  string
	value     float64
	symbol    string
	createdAt time.Time
}

func NewSensorPatientMetric(
	sensorID string,
	patientID string,
	metricID string,
	value float64,
	symbol string,
	createdAt time.Time,
) *SensorPatientMetric {
	return &SensorPatientMetric{
		sensorID:  sensorID,
		patientID: patientID,
		metricID:  metricID,
		value:     value,
		symbol:    symbol,
		createdAt: createdAt,
	}
}

func ReconstituteSensorPatientMetric(p SensorPatientMetricProps) *SensorPatientMetric {
	return &SensorPatientMetric{
		sensorID:  p.SensorID,
		patientID: p.PatientID,
		metricID:  p.MetricID,
		value:     p.Value,
		symbol:    p.Symbol,
		createdAt: p.CreatedAt,
	}
}

func (spm *SensorPatientMetric) SensorID() string     { return spm.sensorID }
func (spm *SensorPatientMetric) PatientID() string    { return spm.patientID }
func (spm *SensorPatientMetric) MetricID() string     { return spm.metricID }
func (spm *SensorPatientMetric) Value() float64       { return spm.value }
func (spm *SensorPatientMetric) Symbol() string       { return spm.symbol }
func (spm *SensorPatientMetric) CreatedAt() time.Time { return spm.createdAt }

func (spm *SensorPatientMetric) SetValue(value float64) *SensorPatientMetric {
	spm.value = value
	return spm
}

func (spm *SensorPatientMetric) SetSymbol(symbol string) *SensorPatientMetric {
	spm.symbol = symbol
	return spm
}
