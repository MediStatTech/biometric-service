package get

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct{}

type Response struct {
	Diseases []domain.DiseaseProps
}
