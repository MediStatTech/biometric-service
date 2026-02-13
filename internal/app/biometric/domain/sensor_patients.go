package domain

import "time"

type SensorPatientStatus string

const (
	SensorPatientStatusActive   SensorPatientStatus = "active"
	SensorPatientStatusInactive SensorPatientStatus = "inactive"
)

func (s SensorPatientStatus) String() string {
	return string(s)
}

type SensorPatientProps struct {
	SensorID  string
	PatientID string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SensorPatient struct {
	sensorID  string
	patientID string
	status    string
	createdAt time.Time
	updatedAt time.Time
}

func NewSensorPatient(
	sensorID string,
	patientID string,
	createdAt time.Time,
) *SensorPatient {
	return &SensorPatient{
		sensorID:  sensorID,
		patientID: patientID,
		status:    SensorPatientStatusActive.String(),
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func ReconstituteSensorPatient(p SensorPatientProps) *SensorPatient {
	return &SensorPatient{
		sensorID:  p.SensorID,
		patientID: p.PatientID,
		status:    p.Status,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (sp *SensorPatient) SensorID() string            { return sp.sensorID }
func (sp *SensorPatient) PatientID() string           { return sp.patientID }
func (sp *SensorPatient) Status() SensorPatientStatus { return SensorPatientStatus(sp.status) }
func (sp *SensorPatient) CreatedAt() time.Time        { return sp.createdAt }
func (sp *SensorPatient) UpdatedAt() time.Time        { return sp.updatedAt }

func (sp *SensorPatient) SetStatus(status SensorPatientStatus) *SensorPatient {
	sp.status = status.String()
	return sp
}

func (sp *SensorPatient) SetActiveStatus() *SensorPatient {
	sp.status = SensorPatientStatusActive.String()
	return sp
}

func (sp *SensorPatient) SetInactiveStatus() *SensorPatient {
	sp.status = SensorPatientStatusInactive.String()
	return sp
}

func (sp *SensorPatient) SetUpdatedAt(updatedAt time.Time) *SensorPatient {
	sp.updatedAt = updatedAt
	return sp
}
