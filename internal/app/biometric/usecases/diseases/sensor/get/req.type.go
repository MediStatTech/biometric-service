package get

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct {
	DiseaseID string
}

type Response struct {
	DiseaseSensors []domain.DiseaseSensorProps
}
