package get

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct{}

type SensorWithMetricTypes struct {
	Sensor      domain.SensorProps
	MetricTypes []domain.MetricTypeProps
}

type Response struct {
	Sensors []SensorWithMetricTypes
}
