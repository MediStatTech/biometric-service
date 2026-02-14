package get

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct {
	SensorID  string
	PatientID string
}

type Response struct {
	Metrics []domain.SensorPatientMetricProps
}
