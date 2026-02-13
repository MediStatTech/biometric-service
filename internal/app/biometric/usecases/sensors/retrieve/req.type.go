package retrieve

import "github.com/MediStatTech/biometric-service/internal/app/biometric/domain"

type Request struct {
	SensorID string
}

type Response struct {
	Sensor domain.SensorProps
}
