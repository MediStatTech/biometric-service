package domain

import (
	"time"

	"github.com/google/uuid"
)

type SensorMetricProps struct {
	SensorID  string
	MetricID  string
	Value     float64
	CreatedAt time.Time
}

type SensorMetric struct {
	sensorID  string
	metricID  string
	value     float64
	createdAt time.Time
}

func NewSensorMetric(
	sensorID string,
	value float64,
	createdAt time.Time,
) *SensorMetric {
	return &SensorMetric{
		sensorID:  sensorID,
		metricID:  uuid.NewString(),
		value:     value,
		createdAt: createdAt,
	}
}

func ReconstituteSensorMetric(p SensorMetricProps) *SensorMetric {
	return &SensorMetric{
		sensorID:  p.SensorID,
		metricID:  p.MetricID,
		value:     p.Value,
		createdAt: p.CreatedAt,
	}
}

func (sm *SensorMetric) SensorID() string     { return sm.sensorID }
func (sm *SensorMetric) MetricID() string     { return sm.metricID }
func (sm *SensorMetric) Value() float64       { return sm.value }
func (sm *SensorMetric) CreatedAt() time.Time { return sm.createdAt }

func (sm *SensorMetric) SetValue(value float64) *SensorMetric {
	sm.value = value
	return sm
}
