package domain

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================================
// Sensor
// ============================================================================

type SensorProps struct {
	SensorID  string
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Sensor struct {
	sensorID  string
	name      string
	code      string
	createdAt time.Time
	updatedAt time.Time
}

func NewSensor(
	name string,
	code string,
	createdAt time.Time,
) *Sensor {
	return &Sensor{
		sensorID:  uuid.NewString(),
		name:      name,
		code:      code,
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func ReconstituteSensor(p SensorProps) *Sensor {
	return &Sensor{
		sensorID:  p.SensorID,
		name:      p.Name,
		code:      p.Code,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (s *Sensor) SensorID() string     { return s.sensorID }
func (s *Sensor) Name() string         { return s.name }
func (s *Sensor) Code() string         { return s.code }
func (s *Sensor) CreatedAt() time.Time { return s.createdAt }
func (s *Sensor) UpdatedAt() time.Time { return s.updatedAt }

func (s *Sensor) SetName(name string) *Sensor {
	s.name = name
	return s
}

func (s *Sensor) SetCode(code string) *Sensor {
	s.code = code
	return s
}

func (s *Sensor) SetUpdatedAt(updatedAt time.Time) *Sensor {
	s.updatedAt = updatedAt
	return s
}