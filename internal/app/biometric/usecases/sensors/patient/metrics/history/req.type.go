package history

import (
	"time"

	metric_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/metrics/get"
)

type Request struct {
	SensorID  string
	PatientID string
	StartTime time.Time
	EndTime   time.Time
	Limit     int32
	Offset    int32
}

type Response struct {
	Measurements []metric_get.Measurement
	Total        int64
}
