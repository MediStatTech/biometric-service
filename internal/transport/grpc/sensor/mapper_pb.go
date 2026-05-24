package sensor

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	metric_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/sensors/patient/metrics/get"
	pb_models "github.com/MediStatTech/biometric-client/pb/go/models/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func sensorPropsToPb(sensor domain.SensorProps, metricTypes []domain.MetricTypeProps) *pb_models.Sensor_Read {
	out := &pb_models.Sensor_Read{
		SensorId: sensor.SensorID,
		Name:     sensor.Name,
		Code:     sensor.Code,
		Symbol:   sensor.Symbol,
	}
	if len(metricTypes) > 0 {
		out.MetricTypes = make([]*pb_models.MetricType_Read, 0, len(metricTypes))
		for _, mt := range metricTypes {
			out.MetricTypes = append(out.MetricTypes, metricTypePropsToPb(mt))
		}
	}
	return out
}

func metricTypePropsToPb(mt domain.MetricTypeProps) *pb_models.MetricType_Read {
	return &pb_models.MetricType_Read{
		MetricTypeId: mt.MetricTypeID,
		SensorId:     mt.SensorID,
		Code:         mt.Code,
		Name:         mt.Name,
		Symbol:       mt.Symbol,
		MinValue:     mt.MinValue,
		MaxValue:     mt.MaxValue,
	}
}

func sensorPatientPropsToPb(sp domain.SensorPatientProps) *pb_models.SensorPatient_Read {
	return &pb_models.SensorPatient_Read{
		SensorId:  sp.SensorID,
		PatientId: sp.PatientID,
		Status:    sp.Status,
	}
}

func measurementToPb(m metric_get.Measurement) *pb_models.SensorPatientMetric_Read {
	components := make([]*pb_models.SensorPatientMetric_Component, 0, len(m.Components))
	for _, c := range m.Components {
		components = append(components, &pb_models.SensorPatientMetric_Component{
			MetricTypeId: c.MetricTypeID,
			Code:         c.Code,
			Name:         c.Name,
			Value:        c.Value,
			Symbol:       c.Symbol,
		})
	}
	return &pb_models.SensorPatientMetric_Read{
		SensorId:   m.SensorID,
		PatientId:  m.PatientID,
		CreatedAt:  timestamppb.New(m.CreatedAt),
		Components: components,
	}
}
