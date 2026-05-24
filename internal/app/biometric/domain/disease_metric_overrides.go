package domain

import "time"

type DiseaseMetricOverrideProps struct {
	DiseaseID    string
	MetricTypeID string
	MinValue     float64
	MaxValue     float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type DiseaseMetricOverride struct {
	diseaseID    string
	metricTypeID string
	minValue     float64
	maxValue     float64
	createdAt    time.Time
	updatedAt    time.Time
}

func NewDiseaseMetricOverride(
	diseaseID string,
	metricTypeID string,
	minValue float64,
	maxValue float64,
	createdAt time.Time,
) *DiseaseMetricOverride {
	return &DiseaseMetricOverride{
		diseaseID:    diseaseID,
		metricTypeID: metricTypeID,
		minValue:     minValue,
		maxValue:     maxValue,
		createdAt:    createdAt,
		updatedAt:    createdAt,
	}
}

func ReconstituteDiseaseMetricOverride(p DiseaseMetricOverrideProps) *DiseaseMetricOverride {
	return &DiseaseMetricOverride{
		diseaseID:    p.DiseaseID,
		metricTypeID: p.MetricTypeID,
		minValue:     p.MinValue,
		maxValue:     p.MaxValue,
		createdAt:    p.CreatedAt,
		updatedAt:    p.UpdatedAt,
	}
}

func (o *DiseaseMetricOverride) DiseaseID() string    { return o.diseaseID }
func (o *DiseaseMetricOverride) MetricTypeID() string { return o.metricTypeID }
func (o *DiseaseMetricOverride) MinValue() float64    { return o.minValue }
func (o *DiseaseMetricOverride) MaxValue() float64    { return o.maxValue }
func (o *DiseaseMetricOverride) CreatedAt() time.Time { return o.createdAt }
func (o *DiseaseMetricOverride) UpdatedAt() time.Time { return o.updatedAt }

func (o *DiseaseMetricOverride) SetRange(min, max float64) *DiseaseMetricOverride {
	o.minValue = min
	o.maxValue = max
	return o
}

func (o *DiseaseMetricOverride) SetUpdatedAt(updatedAt time.Time) *DiseaseMetricOverride {
	o.updatedAt = updatedAt
	return o
}
