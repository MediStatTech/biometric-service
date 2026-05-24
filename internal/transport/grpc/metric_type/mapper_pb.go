package metric_type

import (
	"github.com/MediStatTech/biometric-service/internal/app/biometric/domain"
	pb_models "github.com/MediStatTech/biometric-client/pb/go/models/v1"
)

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
