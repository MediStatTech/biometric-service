package domain

import (
	"time"

	"github.com/google/uuid"
)

type DiseaseProps struct {
	DiseaseID string
	Name      string
	UpdatedAt *time.Time
	CreatedAt time.Time
}

type Disease struct {
	diseaseID string
	name      string
	updatedAt *time.Time
	createdAt time.Time
}

func NewDisease(
	name string,
	createdAt time.Time,
) *Disease {
	return &Disease{
		diseaseID: uuid.NewString(),
		name:      name,
		createdAt: createdAt,
	}
}

func ReconstituteDisease(p DiseaseProps) *Disease {
	return &Disease{
		diseaseID: p.DiseaseID,
		name:      p.Name,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (d *Disease) DiseaseID() string    { return d.diseaseID }
func (d *Disease) Name() string         { return d.name }
func (d *Disease) UpdatedAt() *time.Time { return d.updatedAt }
func (d *Disease) CreatedAt() time.Time { return d.createdAt }

func (d *Disease) SetName(name string) *Disease {
	d.name = name
	return d
}

func (d *Disease) SetUpdatedAt(updatedAt time.Time) *Disease {
	d.updatedAt = &updatedAt
	return d
}
