package retrieve

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct {
	DiseaseID string
}

type Response struct {
	Disease domain.DiseaseProps
}
