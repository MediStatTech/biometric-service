package domain

import (
	"time"

	"github.com/google/uuid"
)

type SensorStatus string

const (
	SensorStatusActive   SensorStatus = "active"
	SensorStatusInactive SensorStatus = "inactive"
	SensorStatusError    SensorStatus = "error"
)

func (s SensorStatus) String() string {
	return string(s)
}

type SensorProps struct {
	SensorID  string
	Name      string
	Status    string
	EnumName  string
	UpdatedAt *time.Time
	CreatedAt time.Time
}

type Sensor struct {
	sensorID  string
	name      string
	status    string
	enumName  string
	updatedAt *time.Time
	createdAt time.Time
}

func NewSensor(
	name string,
	enumName string,
	createdAt time.Time,
) *Sensor {
	return &Sensor{
		sensorID:  uuid.NewString(),
		name:      name,
		status:    SensorStatusActive.String(),
		enumName:  enumName,
		createdAt: createdAt,
	}
}

func ReconstituteSensor(p SensorProps) *Sensor {
	return &Sensor{
		sensorID:  p.SensorID,
		name:      p.Name,
		status:    p.Status,
		enumName:  p.EnumName,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (s *Sensor) SensorID() string       { return s.sensorID }
func (s *Sensor) Name() string           { return s.name }
func (s *Sensor) Status() SensorStatus   { return SensorStatus(s.status) }
func (s *Sensor) EnumName() string       { return s.enumName }
func (s *Sensor) UpdatedAt() *time.Time  { return s.updatedAt }
func (s *Sensor) CreatedAt() time.Time   { return s.createdAt }

func (s *Sensor) SetName(name string) *Sensor {
	s.name = name
	return s
}

func (s *Sensor) SetStatus(status SensorStatus) *Sensor {
	s.status = status.String()
	return s
}

func (s *Sensor) SetActiveStatus() *Sensor {
	s.status = SensorStatusActive.String()
	return s
}

func (s *Sensor) SetInactiveStatus() *Sensor {
	s.status = SensorStatusInactive.String()
	return s
}

func (s *Sensor) SetErrorStatus() *Sensor {
	s.status = SensorStatusError.String()
	return s
}

func (s *Sensor) SetEnumName(enumName string) *Sensor {
	s.enumName = enumName
	return s
}

func (s *Sensor) SetUpdatedAt(updatedAt time.Time) *Sensor {
	s.updatedAt = &updatedAt
	return s
}
