package diseas

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	pb_models "github.com/MediStatTech/biometric-client/pb/go/models/v1"
)

func diseasePropsToPb(disease domain.DiseaseProps) *pb_models.Diseas_Read {
	return &pb_models.Diseas_Read{
		DiseasId: disease.DiseaseID,
		Name:     disease.Name,
		Code:     disease.Code,
	}
}

func diseaseSensorPropsToPb(diseaseSensor domain.DiseaseSensorProps) *pb_models.DiseasSensor_Read {
	return &pb_models.DiseasSensor_Read{
		DiseasId: diseaseSensor.DiseaseID,
		SensorId: diseaseSensor.SensorID,
	}
}
