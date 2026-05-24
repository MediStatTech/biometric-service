package get

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct {
	SensorID string
}

type Response struct {
	MetricTypes []domain.MetricTypeProps
}
