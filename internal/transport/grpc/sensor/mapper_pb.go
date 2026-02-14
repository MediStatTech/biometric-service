package sensor

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	pb_models "github.com/MediStatTech/biometric-client/pb/go/models/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func sensorPropsToPb(sensor domain.SensorProps) *pb_models.Sensor_Read {
	return &pb_models.Sensor_Read{
		SensorId: sensor.SensorID,
		Name:     sensor.Name,
		Code:     sensor.Code,
	}
}

// sensorPatientPropsToPb converts domain.SensorPatientProps to protobuf SensorPatient_Read
func sensorPatientPropsToPb(sensorPatient domain.SensorPatientProps) *pb_models.SensorPatient_Read {
	return &pb_models.SensorPatient_Read{
		SensorId:  sensorPatient.SensorID,
		PatientId: sensorPatient.PatientID,
		Status:    sensorPatient.Status,
	}
}

// sensorPatientMetricPropsToPb converts domain.SensorPatientMetricProps to protobuf SensorPatientMetric_Read
func sensorPatientMetricPropsToPb(metric domain.SensorPatientMetricProps) *pb_models.SensorPatientMetric_Read {
	return &pb_models.SensorPatientMetric_Read{
		SensorId:  metric.SensorID,
		PatientId: metric.PatientID,
		Value:     metric.Value,
		Symbol:    metric.Symbol,
		CreatedAt: timestamppb.New(metric.CreatedAt),
	}
}
