package retrieve

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct {
	DiseaseID string
	SensorID  string
}

type Response struct {
	DiseaseSensor domain.DiseaseSensorProps
}
