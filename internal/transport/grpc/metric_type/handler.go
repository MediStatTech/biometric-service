package metric_type

import (
	"context"

	mt_create "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/metric_types/create"
	mt_get "github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/metric_types/get"
	pb_models "github.com/MediStatTech/biometric-client/pb/go/models/v1"
	pb_services "github.com/MediStatTech/biometric-client/pb/go/services/v1"
)

func (h *Handler) MetricTypeCreate(
	ctx context.Context,
	req *pb_services.MetricTypeCreateRequest,
) (*pb_services.MetricTypeCreateReply, error) {
	if req == nil || req.MetricType == nil {
		return nil, errRequestNil
	}
	if req.SensorId == "" {
		return nil, errInvalidMetricTypeData
	}

	resp, err := h.commands.MetricTypeCreate.Execute(ctx, mt_create.Request{
		SensorID: req.SensorId,
		Code:     req.MetricType.Code,
		Name:     req.MetricType.Name,
		Symbol:   req.MetricType.Symbol,
		MinValue: req.MetricType.MinValue,
		MaxValue: req.MetricType.MaxValue,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to create metric type: %v", err)
		return nil, err
	}

	return &pb_services.MetricTypeCreateReply{
		MetricType: &pb_models.MetricType_Read{
			MetricTypeId: resp.MetricTypeID,
			SensorId:     req.SensorId,
			Code:         req.MetricType.Code,
			Name:         req.MetricType.Name,
			Symbol:       req.MetricType.Symbol,
			MinValue:     req.MetricType.MinValue,
			MaxValue:     req.MetricType.MaxValue,
		},
	}, nil
}

func (h *Handler) MetricTypeGet(
	ctx context.Context,
	req *pb_services.MetricTypeGetRequest,
) (*pb_services.MetricTypeGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.MetricTypeGet.Execute(ctx, mt_get.Request{
		SensorID: req.SensorId,
	})
	if err != nil {
		h.pkg.Logger.Errorf("Failed to get metric types: %v", err)
		return nil, err
	}

	if len(resp.MetricTypes) == 0 {
		return &pb_services.MetricTypeGetReply{
			MetricTypes: []*pb_models.MetricType_Read{},
		}, nil
	}

	out := make([]*pb_models.MetricType_Read, 0, len(resp.MetricTypes))
	for _, mt := range resp.MetricTypes {
		out = append(out, metricTypePropsToPb(mt))
	}
	return &pb_services.MetricTypeGetReply{
		MetricTypes: out,
	}, nil
}
