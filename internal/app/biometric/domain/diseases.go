package domain

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================================
// Disease
// ============================================================================

type DiseaseProps struct {
	DiseaseID string
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Disease struct {
	diseaseID string
	name      string
	code      string
	createdAt time.Time
	updatedAt time.Time
}

func NewDisease(
	name string,
	code string,
	createdAt time.Time,
) *Disease {
	return &Disease{
		diseaseID: uuid.NewString(),
		name:      name,
		code:      code,
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func ReconstituteDisease(p DiseaseProps) *Disease {
	return &Disease{
		diseaseID: p.DiseaseID,
		name:      p.Name,
		code:      p.Code,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (d *Disease) DiseaseID() string     { return d.diseaseID }
func (d *Disease) Name() string          { return d.name }
func (d *Disease) Code() string          { return d.code }
func (d *Disease) CreatedAt() time.Time  { return d.createdAt }
func (d *Disease) UpdatedAt() time.Time  { return d.updatedAt }

func (d *Disease) SetName(name string) *Disease {
	d.name = name
	return d
}

func (d *Disease) SetCode(code string) *Disease {
	d.code = code
	return d
}

func (d *Disease) SetUpdatedAt(updatedAt time.Time) *Disease {
	d.updatedAt = updatedAt
	return d
}

// ============================================================================
// DiseaseSensor
// ============================================================================

type DiseaseSensorProps struct {
	DiseaseID string
	SensorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DiseaseSensor struct {
	diseaseID string
	sensorID  string
	createdAt time.Time
	updatedAt time.Time
}

func NewDiseaseSensor(
	diseaseID string,
	sensorID string,
	createdAt time.Time,
) *DiseaseSensor {
	return &DiseaseSensor{
		diseaseID: diseaseID,
		sensorID:  sensorID,
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func ReconstituteDiseaseSensor(p DiseaseSensorProps) *DiseaseSensor {
	return &DiseaseSensor{
		diseaseID: p.DiseaseID,
		sensorID:  p.SensorID,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (ds *DiseaseSensor) DiseaseID() string    { return ds.diseaseID }
func (ds *DiseaseSensor) SensorID() string     { return ds.sensorID }
func (ds *DiseaseSensor) CreatedAt() time.Time { return ds.createdAt }
func (ds *DiseaseSensor) UpdatedAt() time.Time { return ds.updatedAt }

func (ds *DiseaseSensor) SetUpdatedAt(updatedAt time.Time) *DiseaseSensor {
	ds.updatedAt = updatedAt
	return ds
}
