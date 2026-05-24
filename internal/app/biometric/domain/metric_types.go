package domain

import (
	"time"

	"github.com/google/uuid"
)

type MetricTypeProps struct {
	MetricTypeID string
	SensorID     string
	Code         string
	Name         string
	Symbol       string
	MinValue     float64
	MaxValue     float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type MetricType struct {
	metricTypeID string
	sensorID     string
	code         string
	name         string
	symbol       string
	minValue     float64
	maxValue     float64
	createdAt    time.Time
	updatedAt    time.Time
}

func NewMetricType(
	sensorID string,
	code string,
	name string,
	symbol string,
	minValue float64,
	maxValue float64,
	createdAt time.Time,
) *MetricType {
	return &MetricType{
		metricTypeID: uuid.NewString(),
		sensorID:     sensorID,
		code:         code,
		name:         name,
		symbol:       symbol,
		minValue:     minValue,
		maxValue:     maxValue,
		createdAt:    createdAt,
		updatedAt:    createdAt,
	}
}

func ReconstituteMetricType(p MetricTypeProps) *MetricType {
	return &MetricType{
		metricTypeID: p.MetricTypeID,
		sensorID:     p.SensorID,
		code:         p.Code,
		name:         p.Name,
		symbol:       p.Symbol,
		minValue:     p.MinValue,
		maxValue:     p.MaxValue,
		createdAt:    p.CreatedAt,
		updatedAt:    p.UpdatedAt,
	}
}

func (m *MetricType) MetricTypeID() string { return m.metricTypeID }
func (m *MetricType) SensorID() string     { return m.sensorID }
func (m *MetricType) Code() string         { return m.code }
func (m *MetricType) Name() string         { return m.name }
func (m *MetricType) Symbol() string       { return m.symbol }
func (m *MetricType) MinValue() float64    { return m.minValue }
func (m *MetricType) MaxValue() float64    { return m.maxValue }
func (m *MetricType) CreatedAt() time.Time { return m.createdAt }
func (m *MetricType) UpdatedAt() time.Time { return m.updatedAt }

func (m *MetricType) SetCode(code string) *MetricType {
	m.code = code
	return m
}

func (m *MetricType) SetName(name string) *MetricType {
	m.name = name
	return m
}

func (m *MetricType) SetSymbol(symbol string) *MetricType {
	m.symbol = symbol
	return m
}

func (m *MetricType) SetRange(min, max float64) *MetricType {
	m.minValue = min
	m.maxValue = max
	return m
}

func (m *MetricType) SetUpdatedAt(updatedAt time.Time) *MetricType {
	m.updatedAt = updatedAt
	return m
}
