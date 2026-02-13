package get

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct {
	SensorID  *string
	PatientID *string
	Status    *string
}

type Response struct {
	SensorPatients []domain.SensorPatientProps
}
